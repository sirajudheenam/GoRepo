package main

import (
	"fmt"
	// exitpackage "github.com/sirajudheenam/GoRepo/gobyexample/pkg/exit_package"
	// "github.com/sirajudheenam/GoRepo/gobyexample/pkg/signals"
	// execingprocess "github.com/sirajudheenam/GoRepo/gobyexample/pkg/execing_process"
	// spawningprocess "github.com/sirajudheenam/GoRepo/gobyexample/pkg/spawning_process"
	// contextpackage "github.com/sirajudheenam/GoRepo/gobyexample/pkg/context_package"
	// httpclient "github.com/sirajudheenam/GoRepo/gobyexample/pkg/http_client"
	// httpserver "github.com/sirajudheenam/GoRepo/gobyexample/pkg/http_server"
	// "github.com/sirajudheenam/GoRepo/gobyexample/pkg/logging"
	// envvariables "github.com/sirajudheenam/GoRepo/gobyexample/pkg/env_variables"
	// cmdlinesubcommands "github.com/sirajudheenam/GoRepo/gobyexample/pkg/cmd_line_sub_commands"
	// cmdlineflags "github.com/sirajudheenam/GoRepo/gobyexample/pkg/cmd_line_flags"
	// cmdlinearguments "github.com/sirajudheenam/GoRepo/gobyexample/pkg/cmd_line_arguments"
	// embeddirective "github.com/sirajudheenam/GoRepo/gobyexample/pkg/embed_directive"
	// tempfilesanddirs "github.com/sirajudheenam/GoRepo/gobyexample/pkg/temp_files_and_dirs"
	// "github.com/sirajudheenam/GoRepo/gobyexample/pkg/directories"
	// filepaths "github.com/sirajudheenam/GoRepo/gobyexample/pkg/file_paths"
	// linefilters "github.com/sirajudheenam/GoRepo/gobyexample/pkg/line_filters"
	// writingfiles "github.com/sirajudheenam/GoRepo/gobyexample/pkg/writing_files"
	// readingfiles "github.com/sirajudheenam/GoRepo/gobyexample/pkg/reading_files"
	// base64encoding "github.com/sirajudheenam/GoRepo/gobyexample/pkg/base64_encoding"
	// sha256hashes "github.com/sirajudheenam/GoRepo/gobyexample/pkg/sha256_hashes"
	// urlparsing "github.com/sirajudheenam/GoRepo/gobyexample/pkg/url_parsing"
	// numberparsing "github.com/sirajudheenam/GoRepo/gobyexample/pkg/number_parsing"
	// randomnumbers "github.com/sirajudheenam/GoRepo/gobyexample/pkg/random_numbers"
	// timeformatingparsing "github.com/sirajudheenam/GoRepo/gobyexample/pkg/time_formating_parsing"
	// epochpackage "github.com/sirajudheenam/GoRepo/gobyexample/pkg/epoch_package"
	// timepackage "github.com/sirajudheenam/GoRepo/gobyexample/pkg/time_package"
	// xmlpackage "github.com/sirajudheenam/GoRepo/gobyexample/pkg/xml_package"
	// jsonpackage "github.com/sirajudheenam/GoRepo/gobyexample/pkg/json_package"
	// regularexpressions "github.com/sirajudheenam/GoRepo/gobyexample/pkg/regular_expressions"
	// texttemplates "github.com/sirajudheenam/GoRepo/gobyexample/pkg/text_templates"
	// stringformatting "github.com/sirajudheenam/GoRepo/gobyexample/pkg/string_formatting"
	// stringfunctions "github.com/sirajudheenam/GoRepo/gobyexample/pkg/string_functions"
	// recoverfunction "github.com/sirajudheenam/GoRepo/gobyexample/pkg/recover_function"
	// deferfunction "github.com/sirajudheenam/GoRepo/gobyexample/pkg/defer_function"
	// sortingbyfunctions "github.com/sirajudheenam/GoRepo/gobyexample/pkg/sorting_by_functions"
	// "github.com/sirajudheenam/GoRepo/gobyexample/pkg/sorting"
	// statefulgoroutines "github.com/sirajudheenam/GoRepo/gobyexample/pkg/stateful_go_routines"
	// "github.com/sirajudheenam/GoRepo/gobyexample/pkg/mutexes"
	// atomiccounters "github.com/sirajudheenam/GoRepo/gobyexample/pkg/atomic_counters"
	// ratelimiting "github.com/sirajudheenam/GoRepo/gobyexample/pkg/rate_limiting"
	// waitgroups "github.com/sirajudheenam/GoRepo/gobyexample/pkg/wait_groups"
	// workerpools "github.com/sirajudheenam/GoRepo/gobyexample/pkg/worker_pools"
	// "github.com/sirajudheenam/GoRepo/gobyexample/pkg/tickers"
	// "github.com/sirajudheenam/GoRepo/gobyexample/pkg/arrays"
	// channelbuffering "github.com/sirajudheenam/GoRepo/gobyexample/pkg/channel_buffering"
	// channeldirections "github.com/sirajudheenam/GoRepo/gobyexample/pkg/channel_directions"
	// channelsync "github.com/sirajudheenam/GoRepo/gobyexample/pkg/channel_sync"
	// "github.com/sirajudheenam/GoRepo/gobyexample/pkg/channels"
	// closingchannels "github.com/sirajudheenam/GoRepo/gobyexample/pkg/closing_channels"
	// "github.com/sirajudheenam/GoRepo/gobyexample/pkg/closures"
	// "github.com/sirajudheenam/GoRepo/gobyexample/pkg/constants"
	// customerrors "github.com/sirajudheenam/GoRepo/gobyexample/pkg/custom_errors"
	// "github.com/sirajudheenam/GoRepo/gobyexample/pkg/enums"
	// "github.com/sirajudheenam/GoRepo/gobyexample/pkg/errors"
	// forloop "github.com/sirajudheenam/GoRepo/gobyexample/pkg/for_loop"
	// "github.com/sirajudheenam/GoRepo/gobyexample/pkg/functions"
	// "github.com/sirajudheenam/GoRepo/gobyexample/pkg/generics"
	// goroutines "github.com/sirajudheenam/GoRepo/gobyexample/pkg/go_routines"
	// "github.com/sirajudheenam/GoRepo/gobyexample/pkg/ifelsecond"
	// "github.com/sirajudheenam/GoRepo/gobyexample/pkg/interfaces"
	// "github.com/sirajudheenam/GoRepo/gobyexample/pkg/maps"
	// "github.com/sirajudheenam/GoRepo/gobyexample/pkg/methods"
	// multiplereturnvalues "github.com/sirajudheenam/GoRepo/gobyexample/pkg/multiple_return_values"
	// nonblockingchanneloperations "github.com/sirajudheenam/GoRepo/gobyexample/pkg/non_blocking_channel_operations"
	// "github.com/sirajudheenam/GoRepo/gobyexample/pkg/pointers"
	// rangeoverbuildintypes "github.com/sirajudheenam/GoRepo/gobyexample/pkg/range_over_build_in_types"
	// rangeoverchannels "github.com/sirajudheenam/GoRepo/gobyexample/pkg/range_over_channels"
	// rangeoveriterators "github.com/sirajudheenam/GoRepo/gobyexample/pkg/range_over_iterators"
	// "github.com/sirajudheenam/GoRepo/gobyexample/pkg/recursion"
	// selectcase "github.com/sirajudheenam/GoRepo/gobyexample/pkg/select_case"
	// "github.com/sirajudheenam/GoRepo/gobyexample/pkg/slices"
	// stringsandrunes "github.com/sirajudheenam/GoRepo/gobyexample/pkg/strings_and_runes"
	// structembedding "github.com/sirajudheenam/GoRepo/gobyexample/pkg/struct_embedding"
	// structs "github.com/sirajudheenam/GoRepo/gobyexample/pkg/structs"
	// "github.com/sirajudheenam/GoRepo/gobyexample/pkg/switchcase"
	// "github.com/sirajudheenam/GoRepo/gobyexample/pkg/timeouts"
	// "github.com/sirajudheenam/GoRepo/gobyexample/pkg/timers"
	// "github.com/sirajudheenam/GoRepo/gobyexample/pkg/value"
	// "github.com/sirajudheenam/GoRepo/gobyexample/pkg/variables"
	// variadicfunctions "github.com/sirajudheenam/GoRepo/gobyexample/pkg/variadic_functions"
)

func main() {
	fmt.Println("Hello, Go by Example!")
	// value.Run()
	// multiplereturnvalues.Run()
	// variables.Run()
	// constants.Run()
	// forloop.Run()
	// ifelsecond.Run()
	// switchcase.Run()
	// functions.Run()
	// arrays.Run()
	// slices.Run()
	// maps.Run()
	// variadicfunctions.Run()
	// closures.Run()
	// recursion.Run()
	// rangeoverbuildintypes.Run()
	// pointers.Run()
	// stringsandrunes.Run()
	// structs.Run()
	// methods.Run()
	// interfaces.Run()
	// enums.Run()
	// structembedding.Run()
	// generics.Run()
	// rangeoveriterators.Run()
	// errors.Run()
	// customerrors.Run()
	// goroutines.Run()
	// channels.Run()
	// channelbuffering.Run()
	// channelsync.Run()
	// channeldirections.Run()
	// selectcase.Run()
	// timeouts.Run()
	// nonblockingchanneloperations.Run()
	// closingchannels.Run()
	// rangeoverchannels.Run()
	// timers.Run()
	// tickers.Run()
	// workerpools.Run()
	// waitgroups.Run()
	// ratelimiting.Run()
	// atomiccounters.Run()
	// mutexes.Run()
	// statefulgoroutines.Run()
	// sorting.Run()
	// sortingbyfunctions.Run()
	// deferfunction.Run()
	// recoverfunction.Run()
	// stringfunctions.Run()
	// stringformatting.Run()
	// texttemplates.Run()
	// regularexpressions.Run()
	// jsonpackage.Run()
	// xmlpackage.Run()
	// timepackage.Run()
	// epochpackage.Run()
	// timeformatingparsing.Run()
	// randomnumbers.Run()
	// numberparsing.Run()
	// urlparsing.Run()
	// sha256hashes.Run()
	// base64encoding.Run()

	// // echo "hello" > /tmp/dat
	// // echo "go" >>   /tmp/dat
	// // go run main.go
	// readingfiles.Run()

	// writingfiles.Run()

	// // echo 'hello'   > /tmp/lines
	// // echo 'filter' >> /tmp/lines
	// // cat /tmp/lines | go run main.go
	// linefilters.Run()

	// filepaths.Run()

	// directories.Run()
	// tempfilesanddirs.Run()

	// // $ mkdir -p folder
	// // $ echo "hello go" > folder/single_file.txt
	// // $ echo "123" > folder/file1.hash
	// // $ echo "456" > folder/file2.hash
	// embeddirective.Run()

	// // cd testing_benchmarking
	// // go test -v

	// // go build command-line-arguments.go
	// // $ ./command-line-arguments a b c d
	// // [./command-line-arguments a b c d]

	// // go run main.go 1 2 3 4 5
	// cmdlinearguments.Run()

	// // go run main.go -word=opt -numb=7 -fork -svar=flag
	// // go run main.go -word=opt
	// // go run main.go -word=opt a1 a2 a3
	// // go run main.go -word=opt a1 a2 a3 -numb=7
	// // go run main.go -h
	// // go run main.go -wat
	// cmdlineflags.Run()

	// // go run main.go foo -enable -name=joe a1 a2
	// // go run main.go bar -level 8 a1
	// // go run main.go bar -enable a1
	// cmdlinesubcommands.Run()

	// // BAR=1 go run main.go
	// envvariables.Run()

	// logging.Run()

	// httpclient.Run()

	// Access it in diff terminal curl localhost:8090/hello
	// httpserver.Run()

	// Access it in diff terminal curl localhost:8090/hello
	// contextpackage.Run()

	// spawningprocess.Run()

	// execingprocess.Run()

	// // Press Ctl + C to stop
	// signals.Run()

	// exitpackage.Run()

	// panicfunction.Run()
}
