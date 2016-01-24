package main

import (
	"github.com/codegangsta/cli"
	"github.com/gi4nks/quant"
	"os"
)

var parrot = quant.NewParrot("decanter")

var settings = Settings{}
var repository = Repository{}

func initDB() {
	repository.InitDB(settings.RepositoryUrl())
}

func readSettings() {
	settings.LoadSettings()

	if settings.DebugMode() {
		parrot = quant.NewVerboseParrot("decanter")
	}

	parrot.Debug("Parrot is set to talk so much!")
}

func main() {
	readSettings()
	initDB()

	// -------------------
	app := cli.NewApp()
	app.Name = "decanter"
	app.Usage = "Voxxed Zurich Votes translator"
	app.Version = "0.0.1"
	app.Copyright = "gi4nks - 2015"

	app.Commands = []cli.Command{
		{
			Name:    "run",
			Aliases: []string{"ru"},
			Usage:   "run a command, remember to run the command with -- before. (./ambros r -- ls -la)",
			Action:  CmdRun,
		},
		{
			Name:    "reset",
			Aliases: []string{"re"},
			Usage:   "reset votes",
			Action:  CmdReset,
		},
		{
			Name:    "mock",
			Aliases: []string{"mo"},
			Usage:   "mock server",
			Action:  CmdMock,
		},
	}

	app.Run(os.Args)

	//defer closeDB()
}

// List of functions
func CmdReset(ctx *cli.Context) {
	commandWrapper(ctx, func() {
		parrot.Info("==> Recovery of all booked items not shipped.")

		err := repository.ResetVotes()
		if err != nil {
			parrot.Error("Something went wrong. Please check the error.", err)
			panic(err)
		}
	})
}

func CmdRun(ctx *cli.Context) {
	commandWrapper(ctx, func() {
		parrot.Info("Running decanter...")

		terminated := false

		repository.ResetVotes()

		var rc = RestClient{}

		for !terminated {
			vote := repository.GetNextVote()

			if vote.IsEmpty() {
				terminated = true
			} else {
				// Now I have a vote to send
				err := rc.Send(vote)
				if err != nil {
					parrot.Error("Something went wrong. Please check the error.", err)
					panic(err)
				}

				err = repository.ArchiveVote(vote.ID)
				if err != nil {
					parrot.Error("Something went wrong. Please check the error.", err)
					panic(err)
				}
			}
		}

	})
}

func CmdMock(ctx *cli.Context) {
	commandWrapper(ctx, func() {
		parrot.Info("==> Mocking a server on port 8080")
		// url -i -X POST -H "Content-Type: application/json" -d "{ \"talkid\": \"Talk1\", \"rating\": \"2\" }" http://localhost:8080/api/v1/votes
		serve()
	})
}

func CmdWrapper(ctx *cli.Context) {
}

// -------------------------------
// Cli command wrapper
// -------------------------------
func commandWrapper(ctx *cli.Context, cmd quant.Action0) {
	CmdWrapper(ctx)

	cmd()
}
