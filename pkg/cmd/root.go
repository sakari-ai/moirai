package cmd

import (
	"context"
	"fmt"
	"github.com/sakari-ai/moirai/bootstrap"
	cache "github.com/sakari-ai/moirai/cache/arc"
	"github.com/sakari-ai/moirai/cmd"
	"github.com/sakari-ai/moirai/config"
	"github.com/sakari-ai/moirai/config/loader"
	"github.com/sakari-ai/moirai/log"
	"github.com/sakari-ai/moirai/pkg/handler"
	"github.com/sakari-ai/moirai/pkg/model"
	"github.com/sakari-ai/moirai/pkg/storage"
	"github.com/sakari-ai/moirai/proto"
	"github.com/sakari-ai/moirai/server"
	"github.com/sakari-ai/moirai/server/grpc"
	"github.com/sakari-ai/moirai/server/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	googlegrpc "google.golang.org/grpc"
)

var aceCmd = &cobra.Command{
	Use:   "moirai",
	Short: "Start MOIRAI Service",
	Long:  "OK babe",
	Run: func(cmd *cobra.Command, args []string) {
		startAce()
	},
}

var (
	onClose func()
)

func init() {
	cmd.Register(aceCmd)

	aceCmd.PersistentFlags().String("namespace", "ace", "config namespace")
	viper.BindPFlag("namespace", aceCmd.PersistentFlags().Lookup("namespace"))
}

func startAce() {
	var (
		s             = server.Server{}
		mr            = new(handler.Moirai)
		registerAsset = registerMoirai(mr)
		ctx, cancel   = context.WithCancel(context.Background())
	)
	defer cancel()

	var objects = []bootstrap.Object{
		bootstrap.ByValue(&s),
		//moirai
		bootstrap.ByName("moirai", mr),
	}

	if err := bootstrap.Populate(func(loader loader.Loader) []bootstrap.Object {
		cfg := &config.Config{}
		config.LoadWithPlaceholder(loader, cfg)
		bootstrap.LoadDB(cfg.Database)
		grpcCfg, httpCfg := cfg.GRPC, cfg.HTTP
		objects = append(objects,
			bootstrap.ByName("db_cache", cache.New(ctx, 100)),
			bootstrap.ByName("schema_storage", storage.NewStorage()),
			bootstrap.ByName("schema_validator", model.NewValidator()),
			bootstrap.ByName("grpc", newGRPCServer(grpcServerOption{
				RegisterFuncs: []grpc.RegisterFunc{registerAsset},
				Config:        grpcCfg,
			})),
			bootstrap.ByName("http", newHTTPServer(httpServerOption{
				ConfigGRPC: grpcCfg,
				ConfigHTTP: httpCfg,
			})),
		)
		return objects
	}); err != nil {
		panic(err)
	}

	if err := s.Run(ctx); err != nil {
		panic(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	select {
	case s := <-c:
		log.Info(fmt.Sprintf("received signal: %s", s))
	case <-ctx.Done():
		onClose()
		log.Info("onClose received context done")
	}
}

type grpcServerOption struct {
	RegisterFuncs []grpc.RegisterFunc
	Config        *config.GRPC
	JWTSecret     string
}

func newGRPCServer(opt grpcServerOption) *grpc.Server {
	s := grpc.New(
		grpc.WithAddress(opt.Config.Address),
		grpc.WithRegisterFunc(opt.RegisterFuncs...),
	)
	return s
}

type httpServerOption struct {
	ConfigGRPC *config.GRPC
	ConfigHTTP *config.HTTP
}

func newHTTPServer(opt httpServerOption) *http.Server {
	s := http.New(
		http.WithGRPCAddress(opt.ConfigGRPC.Address),
		http.WithAddress(opt.ConfigHTTP.Address),
		http.WithHandlerFunc(proto.RegisterBackendHandlerFromEndpoint),
	)

	return s
}

func registerMoirai(mr *handler.Moirai) func(*googlegrpc.Server) {
	return func(s *googlegrpc.Server) {
		proto.RegisterBackendServer(s, mr)
	}
}
