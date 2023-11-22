package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	d "github.com/artag/clnr/domain"
	commands "github.com/artag/clnr/implementations/commands"
	queries "github.com/artag/clnr/implementations/queries"
	_ "github.com/artag/clnr/interfaces"
)

const (
	// Application version
	version = "0.0.1"
)

var (
	// Current year
	year int = time.Now().Year()
)

func main() {
	ruleFlag := flag.String("rule", "", "Rule to delete directories or files.")
	fileRuleFlag := flag.Bool("file", false, "True - rule to delete files, false - rule to delete directories.")
	excludeFlag := flag.Bool("exclude", false, "True - exclude directory or file from delete list, false - include directory or file to delete list.")
	equalFlag := flag.String("equal", "", "Include/exclude directory or file by name equality.")
	containFlag := flag.String("contain", "", "Include/exclude directory or file by name containing substring.")
	startFlag := flag.String("start", "", "Include/exclude directory or file by name beginning with prefix.")
	endFlag := flag.String("end", "", "Include/exclude directory or file by name ending with suffix.")
	rootFlag := flag.String("root", "", "Add a root directory to search for directories or files to delete.")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "clnr (Cleaner) ver. %s.\n", version)
		fmt.Fprintln(flag.CommandLine.Output(), "The tool to delete directories and/or files. Use at your own risk.")
		fmt.Fprintf(flag.CommandLine.Output(), "Copyright %d. github.com/artag/clnr\n", year)
		fmt.Fprintln(flag.CommandLine.Output(), "Usage information:")
		flag.PrintDefaults()
	}

	flag.Parse()

	args := d.NewArgs(*ruleFlag, *fileRuleFlag, !*excludeFlag, *equalFlag, *containFlag, *startFlag, *endFlag, *rootFlag, flag.Args())

	program, err := createProgram()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := program.Run(args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		flag.Usage()
		os.Exit(1)
	}
}

func createProgram() (*Program, error) {
	clearConsole, err := commands.NewClearCli()
	if err != nil {
		return nil, err
	}

	printer := commands.NewPrintCli()
	selectOptions := commands.NewSelectOptionCli(clearConsole, printer)
	excludeDirs := commands.NewExcludeDirsCli(clearConsole, printer, selectOptions)
	excludeFiles := commands.NewExcludeFilesCli(clearConsole, printer, selectOptions)

	program := Program{
		queries.NewGetRule(),
		commands.NewSaveRule(),
		queries.NewSearchDirsToDelete(),
		queries.NewSearchFilesToDelete(),
		commands.NewSelectDirsCli(clearConsole, excludeDirs, printer, selectOptions),
		commands.NewSelectFilesCli(clearConsole, excludeFiles, printer, selectOptions),
		commands.NewDeleteDirs(printer),
		commands.NewDeleteFiles(printer),
	}

	return &program, nil
}
