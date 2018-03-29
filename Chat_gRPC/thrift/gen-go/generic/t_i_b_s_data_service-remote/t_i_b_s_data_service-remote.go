// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"math"
	"net"
	"net/url"
	"openstars/core/bigset/generic"
	"os"
	"strconv"
	"strings"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  TPutItemResult putItem(TKey bigsetID, TItem item)")
	fmt.Fprintln(os.Stderr, "  bool removeItem(TKey bigsetID, TItemKey itemKey)")
	fmt.Fprintln(os.Stderr, "  TExistedResult existed(TKey bigsetID, TItemKey itemKey)")
	fmt.Fprintln(os.Stderr, "  TItemResult getItem(TKey bigsetID, TItemKey itemKey)")
	fmt.Fprintln(os.Stderr, "  TItemSetResult getSlice(TKey bigsetID, i32 fromIndex, i32 count)")
	fmt.Fprintln(os.Stderr, "  TItemSetResult getSliceFromItem(TKey bigsetID, TItemKey fromKey, i32 count)")
	fmt.Fprintln(os.Stderr, "  TItemSetResult getSliceR(TKey bigsetID, i32 fromIndex, i32 count)")
	fmt.Fprintln(os.Stderr, "  TItemSetResult getSliceFromItemR(TKey bigsetID, TItemKey fromKey, i32 count)")
	fmt.Fprintln(os.Stderr, "  TItemSetResult rangeQuery(TKey bigsetID, TItemKey startKey, TItemKey endKey)")
	fmt.Fprintln(os.Stderr, "  bool bulkLoad(TKey bigsetID, TItemSet setData)")
	fmt.Fprintln(os.Stderr, "  TMultiPutItemResult multiPut(TKey bigsetID, TItemSet setData, bool getAddedItems, bool getReplacedItems)")
	fmt.Fprintln(os.Stderr, "  i64 getTotalCount(TKey bigsetID)")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	var parsedUrl url.URL
	var trans thrift.TTransport
	_ = strconv.Atoi
	_ = math.Abs
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.Parse()

	if len(urlString) > 0 {
		parsedUrl, err := url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
	} else if useHttp {
		_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
	}

	cmd := flag.Arg(0)
	var err error
	if useHttp {
		trans, err = thrift.NewTHttpClient(parsedUrl.String())
	} else {
		portStr := fmt.Sprint(port)
		if strings.Contains(host, ":") {
			host, portStr, err = net.SplitHostPort(host)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error with host:", err)
				os.Exit(1)
			}
		}
		trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}
		if framed {
			trans = thrift.NewTFramedTransport(trans)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}
	defer trans.Close()
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
		break
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
		break
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
		Usage()
		os.Exit(1)
	}
	client := generic.NewTIBSDataServiceClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "putItem":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "PutItem requires 2 args")
			flag.Usage()
		}
		argvalue0, err211 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err211 != nil {
			Usage()
			return
		}
		value0 := generic.TKey(argvalue0)
		arg212 := flag.Arg(2)
		mbTrans213 := thrift.NewTMemoryBufferLen(len(arg212))
		defer mbTrans213.Close()
		_, err214 := mbTrans213.WriteString(arg212)
		if err214 != nil {
			Usage()
			return
		}
		factory215 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt216 := factory215.GetProtocol(mbTrans213)
		argvalue1 := generic.NewTItem()
		err217 := argvalue1.Read(jsProt216)
		if err217 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.PutItem(value0, value1))
		fmt.Print("\n")
		break
	case "removeItem":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "RemoveItem requires 2 args")
			flag.Usage()
		}
		argvalue0, err218 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err218 != nil {
			Usage()
			return
		}
		value0 := generic.TKey(argvalue0)
		argvalue1 := []byte(flag.Arg(2))
		value1 := generic.TItemKey(argvalue1)
		fmt.Print(client.RemoveItem(value0, value1))
		fmt.Print("\n")
		break
	case "existed":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "Existed requires 2 args")
			flag.Usage()
		}
		argvalue0, err220 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err220 != nil {
			Usage()
			return
		}
		value0 := generic.TKey(argvalue0)
		argvalue1 := []byte(flag.Arg(2))
		value1 := generic.TItemKey(argvalue1)
		fmt.Print(client.Existed(value0, value1))
		fmt.Print("\n")
		break
	case "getItem":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "GetItem requires 2 args")
			flag.Usage()
		}
		argvalue0, err222 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err222 != nil {
			Usage()
			return
		}
		value0 := generic.TKey(argvalue0)
		argvalue1 := []byte(flag.Arg(2))
		value1 := generic.TItemKey(argvalue1)
		fmt.Print(client.GetItem(value0, value1))
		fmt.Print("\n")
		break
	case "getSlice":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "GetSlice requires 3 args")
			flag.Usage()
		}
		argvalue0, err224 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err224 != nil {
			Usage()
			return
		}
		value0 := generic.TKey(argvalue0)
		tmp1, err225 := (strconv.Atoi(flag.Arg(2)))
		if err225 != nil {
			Usage()
			return
		}
		argvalue1 := int32(tmp1)
		value1 := argvalue1
		tmp2, err226 := (strconv.Atoi(flag.Arg(3)))
		if err226 != nil {
			Usage()
			return
		}
		argvalue2 := int32(tmp2)
		value2 := argvalue2
		fmt.Print(client.GetSlice(value0, value1, value2))
		fmt.Print("\n")
		break
	case "getSliceFromItem":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "GetSliceFromItem requires 3 args")
			flag.Usage()
		}
		argvalue0, err227 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err227 != nil {
			Usage()
			return
		}
		value0 := generic.TKey(argvalue0)
		argvalue1 := []byte(flag.Arg(2))
		value1 := generic.TItemKey(argvalue1)
		tmp2, err229 := (strconv.Atoi(flag.Arg(3)))
		if err229 != nil {
			Usage()
			return
		}
		argvalue2 := int32(tmp2)
		value2 := argvalue2
		fmt.Print(client.GetSliceFromItem(value0, value1, value2))
		fmt.Print("\n")
		break
	case "getSliceR":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "GetSliceR requires 3 args")
			flag.Usage()
		}
		argvalue0, err230 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err230 != nil {
			Usage()
			return
		}
		value0 := generic.TKey(argvalue0)
		tmp1, err231 := (strconv.Atoi(flag.Arg(2)))
		if err231 != nil {
			Usage()
			return
		}
		argvalue1 := int32(tmp1)
		value1 := argvalue1
		tmp2, err232 := (strconv.Atoi(flag.Arg(3)))
		if err232 != nil {
			Usage()
			return
		}
		argvalue2 := int32(tmp2)
		value2 := argvalue2
		fmt.Print(client.GetSliceR(value0, value1, value2))
		fmt.Print("\n")
		break
	case "getSliceFromItemR":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "GetSliceFromItemR requires 3 args")
			flag.Usage()
		}
		argvalue0, err233 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err233 != nil {
			Usage()
			return
		}
		value0 := generic.TKey(argvalue0)
		argvalue1 := []byte(flag.Arg(2))
		value1 := generic.TItemKey(argvalue1)
		tmp2, err235 := (strconv.Atoi(flag.Arg(3)))
		if err235 != nil {
			Usage()
			return
		}
		argvalue2 := int32(tmp2)
		value2 := argvalue2
		fmt.Print(client.GetSliceFromItemR(value0, value1, value2))
		fmt.Print("\n")
		break
	case "rangeQuery":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "RangeQuery requires 3 args")
			flag.Usage()
		}
		argvalue0, err236 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err236 != nil {
			Usage()
			return
		}
		value0 := generic.TKey(argvalue0)
		argvalue1 := []byte(flag.Arg(2))
		value1 := generic.TItemKey(argvalue1)
		argvalue2 := []byte(flag.Arg(3))
		value2 := generic.TItemKey(argvalue2)
		fmt.Print(client.RangeQuery(value0, value1, value2))
		fmt.Print("\n")
		break
	case "bulkLoad":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "BulkLoad requires 2 args")
			flag.Usage()
		}
		argvalue0, err239 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err239 != nil {
			Usage()
			return
		}
		value0 := generic.TKey(argvalue0)
		arg240 := flag.Arg(2)
		mbTrans241 := thrift.NewTMemoryBufferLen(len(arg240))
		defer mbTrans241.Close()
		_, err242 := mbTrans241.WriteString(arg240)
		if err242 != nil {
			Usage()
			return
		}
		factory243 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt244 := factory243.GetProtocol(mbTrans241)
		argvalue1 := generic.NewTItemSet()
		err245 := argvalue1.Read(jsProt244)
		if err245 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.BulkLoad(value0, value1))
		fmt.Print("\n")
		break
	case "multiPut":
		if flag.NArg()-1 != 4 {
			fmt.Fprintln(os.Stderr, "MultiPut requires 4 args")
			flag.Usage()
		}
		argvalue0, err246 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err246 != nil {
			Usage()
			return
		}
		value0 := generic.TKey(argvalue0)
		arg247 := flag.Arg(2)
		mbTrans248 := thrift.NewTMemoryBufferLen(len(arg247))
		defer mbTrans248.Close()
		_, err249 := mbTrans248.WriteString(arg247)
		if err249 != nil {
			Usage()
			return
		}
		factory250 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt251 := factory250.GetProtocol(mbTrans248)
		argvalue1 := generic.NewTItemSet()
		err252 := argvalue1.Read(jsProt251)
		if err252 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		argvalue2 := flag.Arg(3) == "true"
		value2 := argvalue2
		argvalue3 := flag.Arg(4) == "true"
		value3 := argvalue3
		fmt.Print(client.MultiPut(value0, value1, value2, value3))
		fmt.Print("\n")
		break
	case "getTotalCount":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetTotalCount requires 1 args")
			flag.Usage()
		}
		argvalue0, err255 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err255 != nil {
			Usage()
			return
		}
		value0 := generic.TKey(argvalue0)
		fmt.Print(client.GetTotalCount(value0))
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}
