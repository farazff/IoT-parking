package serve

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/okian/servo/v2"
	_ "github.com/okian/servo/v2/config"
	_ "github.com/okian/servo/v2/db"
	_ "github.com/okian/servo/v2/lg"
	_ "github.com/okian/servo/v2/rest"
	"github.com/spf13/cobra"
	_ "go.uber.org/automaxprocs"

	_ "github.com/farazff/IoT-parking/io/rest"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT, syscall.SIGABRT, syscall.SIGHUP)
	ctx, cl := context.WithCancel(context.Background())
	final := servo.Initialize(ctx)
	<-sigs
	final()
	cl()
}

// Register API command
func Register() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "parking http api",
		Run: func(cmd *cobra.Command, args []string) {
			main()
		},
	}
}
