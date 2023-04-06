package main

import (
	//UI

	"reflect"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	//"github.com/epiclabs-io/winman"

	//HW Data

	"os"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"

	//Mail

	"fmt"
	"net/smtp"

	//CSV

	//Various
	"bufio"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// Define constants
const (
	CSVDelimiter            = ";" //Used as CSV Delimiter
	FieldDelimiter          = "=" //Used to separate field name and field value (to serialize structure)
	ArrayDelimiter          = "↔" //ASCII 29 - Used as Array delimiter (to serialize array)
	NetInterfacesDelimiter1 = "▲" //ASCII 30 - Used as fields delimiter in NetInterfaces
	NetInterfacesDelimiter2 = "▼" //ASCII 31 - Used as subfields delimiter in NetInterfaces
)

var NewLineChar = "\n"

type S_NetInterface struct {
	//Structure of Net Interface object
	InterfaceName            string
	MACAddress               string
	Interfacebehaviororflags []string
	IPv6_IPv4addresses       []string
}

func (e *S_NetInterface) ToString() string {
	//Convert S_NetInterface to string
	var NewString string
	NewString = e.InterfaceName + NetInterfacesDelimiter1
	NewString += e.MACAddress + NetInterfacesDelimiter1
	NewString += strings.Join(e.Interfacebehaviororflags, NetInterfacesDelimiter2) + NetInterfacesDelimiter1
	NewString += strings.Join(e.IPv6_IPv4addresses, NetInterfacesDelimiter2)
	return NewString
}

func StringToNetInterface(NetInterfacesString string) S_NetInterface {
	//Convert string to S_NetInterface
	var NetInterface S_NetInterface

	TempArray := strings.Split(NetInterfacesString, NetInterfacesDelimiter1)
	NetInterface.InterfaceName = TempArray[0]
	NetInterface.MACAddress = TempArray[2]
	NetInterface.Interfacebehaviororflags = strings.Split(TempArray[3], NetInterfacesDelimiter2)
	NetInterface.IPv6_IPv4addresses = strings.Split(TempArray[4], NetInterfacesDelimiter2)
	return NetInterface
}

type S_HealthSnapshot struct {
	//Structure of Health Snapshot object

	Date string
	Time string

	Hostname string
	Uptime   string

	runtimeOS string
	OS        string
	Platform  string
	HostID    string

	CPUindexNumber string
	CPUVendorID    string
	CPUFamily      string
	CPUModelName   string
	CPUSpeed       string
	CPUNumberCores string

	NumberofProcessesRunning string
	CurrentCPUutilization    []string

	MemoryTotal          string
	MemoryFree           string
	MemoryPercentageUsed string

	DiskSpaceTotal           string
	DiskSpaceUsed            string
	DiskSpaceFree            string
	DiskSpaceUsagePercentage string

	NetInterfaces []S_NetInterface

	Test1 int
	Test2 []int
	Test3 uint
	Test4 bool
	Test5 int64
}

/*
func (e *S_HealthSnapshot) GetField(field string) string {
	r := reflect.ValueOf(e)
	f := reflect.Indirect(r).FieldByName(field)
	return f.Interface().(string)
}
*/

func (e S_HealthSnapshot) ToLogLine() string {
	//Serialize HealthSnapshot to CSV string

	//fmt.Printf(e.DiskSpaceFree)

	var in int      //only to test type. other solution ?
	var i32 int32   //only to test type. other solution ?
	var i64 int64   //only to test type. other solution ?
	var ui uint     //only to test type. other solution ?
	var ui32 uint32 //only to test type. other solution ?
	var ui64 uint64 //only to test type. other solution ?
	var f32 float32 //only to test type. other solution ?
	var f64 float64 //only to test type. other solution ?
	var b bool      //only to test type. other solution ?

	var NewString string

	//v = reflect.Indirect(familyPtr).FieldByName("last")

	fields := reflect.TypeOf(e)
	values := reflect.ValueOf(e)

	num := values.NumField()
	for i := 0; i < num; i++ {
		field := fields.Field(i)
		value := values.Field(i)
		//fmt.Print("Type:", field.Type, ", ", field.Name, "=", value, "  kind=", value.Kind(), "\n")
		switch field.Type {
		case reflect.TypeOf(""):
			NewString += field.Name + FieldDelimiter + value.String() + CSVDelimiter
		case reflect.TypeOf([]string{}):
			var NewArray []string = value.Interface().([]string)
			NewString += field.Name + FieldDelimiter + strings.Join(NewArray, ArrayDelimiter) + CSVDelimiter
		case reflect.TypeOf(in):
			NewString += field.Name + FieldDelimiter + strconv.FormatInt(value.Int(), 10) + CSVDelimiter
		case reflect.TypeOf([]int{}):
			var IntString string = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(value.Interface().([]int))), ArrayDelimiter), "[]")
			//var NewArray []string = value.Interface().([]int)
			NewString += field.Name + FieldDelimiter + IntString + CSVDelimiter
		case reflect.TypeOf(i32):
			NewString += field.Name + FieldDelimiter + strconv.FormatInt(int64(value.Interface().(int32)), 10) + CSVDelimiter
		case reflect.TypeOf(i64):
			NewString += field.Name + FieldDelimiter + strconv.FormatInt(value.Interface().(int64), 10) + CSVDelimiter
		case reflect.TypeOf(ui):
			NewString += field.Name + FieldDelimiter + strconv.FormatUint(value.Uint(), 10) + CSVDelimiter
		case reflect.TypeOf(ui32):
			NewString += field.Name + FieldDelimiter + strconv.FormatUint(value.Uint(), 10) + CSVDelimiter
		case reflect.TypeOf(ui64):
			NewString += field.Name + FieldDelimiter + strconv.FormatUint(value.Interface().(uint64), 10) + CSVDelimiter
		case reflect.TypeOf(f32):
			NewString += field.Name + FieldDelimiter + strconv.FormatFloat(value.Float(), 'e', -1, 32) + CSVDelimiter
		case reflect.TypeOf(f64):
			NewString += field.Name + FieldDelimiter + strconv.FormatFloat(value.Float(), 'e', -1, 64) + CSVDelimiter
		case reflect.TypeOf(b):
			NewString += field.Name + FieldDelimiter + strconv.FormatBool(value.Bool()) + CSVDelimiter
		case reflect.TypeOf([]S_NetInterface{}):
			var NetInterfaces []S_NetInterface = value.Interface().([]S_NetInterface)
			var NetInterfacesString string
			for _, NetInterface := range NetInterfaces {
				NetInterfacesString = NetInterfacesString + NetInterface.ToString() + ArrayDelimiter
			}
			NewString += field.Name + FieldDelimiter + NetInterfacesString + CSVDelimiter
		default:
			dealwithErr(fmt.Errorf("StructurToString: error converting " + field.Name + "(" + field.Type.String() + ")"))
		}
	}

	return NewString
}

func LogLineToHealthSnapshot(csvline string) S_HealthSnapshot {
	//Convert csv string to S_HealthSnapshot

	var NewHealthSnapshot S_HealthSnapshot

	var Fields []string = strings.Split(csvline, CSVDelimiter)

	for _, Field := range Fields {
		MyFields := strings.Split(Field, FieldDelimiter)
		MyFieldName := MyFields[0]
		MyFieldValue := MyFields[1]

		switch MyFieldName {
		case "Date":
			NewHealthSnapshot.Date = MyFieldValue
		case "Time":
			NewHealthSnapshot.Time = MyFieldValue

		case "Hostname":
			NewHealthSnapshot.Hostname = MyFieldValue
		case "Uptime":
			NewHealthSnapshot.Uptime = MyFieldValue
		case "runtimeOS":
			NewHealthSnapshot.runtimeOS = MyFieldValue
		case "OS":
			NewHealthSnapshot.OS = MyFieldValue
		case "Platform":
			NewHealthSnapshot.Platform = MyFieldValue
		case "HostID":
			NewHealthSnapshot.HostID = MyFieldValue

		case "CPUindexNumber":
			NewHealthSnapshot.CPUindexNumber = MyFieldValue
		case "CPUVendorID":
			NewHealthSnapshot.CPUVendorID = MyFieldValue
		case "CPUFamily":
			NewHealthSnapshot.CPUFamily = MyFieldValue
		case "CPUNumberCores":
			NewHealthSnapshot.CPUNumberCores = MyFieldValue
		case "CPUModelName":
			NewHealthSnapshot.CPUModelName = MyFieldValue
		case "CPUSpeed":
			NewHealthSnapshot.CPUSpeed = MyFieldValue

		case "NumberofProcessesRunning":
			NewHealthSnapshot.NumberofProcessesRunning = MyFieldValue
		case "CurrentCPUutilization":
			NewHealthSnapshot.CurrentCPUutilization = strings.Split(MyFieldValue, ArrayDelimiter)

		case "MemoryTotal":
			NewHealthSnapshot.MemoryTotal = MyFieldValue
		case "MemoryFree":
			NewHealthSnapshot.MemoryFree = MyFieldValue
		case "MemoryPercentageUsed":
			NewHealthSnapshot.MemoryPercentageUsed = MyFieldValue

		case "DiskSpaceTotal":
			NewHealthSnapshot.DiskSpaceTotal = MyFieldValue
		case "DiskSpaceUsed":
			NewHealthSnapshot.DiskSpaceUsed = MyFieldValue
		case "DiskSpaceFree":
			NewHealthSnapshot.DiskSpaceFree = MyFieldValue
		case "DiskSpaceUsagePercentage":
			NewHealthSnapshot.DiskSpaceUsagePercentage = MyFieldValue

		case "NetInterfaces":
			MyInterfacesStrings := strings.Split(MyFieldValue, ArrayDelimiter)
			var MyInterfaces []S_NetInterface

			for _, InterfaceString := range MyInterfacesStrings {
				var MyInterface S_NetInterface = StringToNetInterface(InterfaceString)
				MyInterfaces = append(MyInterfaces, MyInterface)
			}
			NewHealthSnapshot.NetInterfaces = MyInterfaces
		default:
		}
	}
	return NewHealthSnapshot
}

func ExportToCSVArray(HealthSnapshots []S_HealthSnapshot) []string {
	//Serialize HealthSnapshot to CSV string

	var in int      //only to test type. other solution ?
	var i32 int32   //only to test type. other solution ?
	var i64 int64   //only to test type. other solution ?
	var ui uint     //only to test type. other solution ?
	var ui32 uint32 //only to test type. other solution ?
	var ui64 uint64 //only to test type. other solution ?
	var f32 float32 //only to test type. other solution ?
	var f64 float64 //only to test type. other solution ?
	var b bool      //only to test type. other solution ?

	//Get titles
	var HSS S_HealthSnapshot
	HSSFields := reflect.TypeOf(HSS)
	HSSnum := HSSFields.NumField()
	var CSVTitles string
	for i := 0; i < HSSnum; i++ {
		field := HSSFields.Field(i)
		CSVTitles += field.Name + CSVDelimiter
	}
	CSVTitles = strings.TrimRight(CSVTitles, CSVDelimiter)

	//New array with all CSV lines
	var NewCSVArray []string
	NewCSVArray = append(NewCSVArray, CSVTitles)

	for _, HealthSnapshot := range HealthSnapshots {

		var NewString string

		//v = reflect.Indirect(familyPtr).FieldByName("last")

		fields := reflect.TypeOf(HealthSnapshot)
		values := reflect.ValueOf(HealthSnapshot)

		num := values.NumField()
		for i := 0; i < num; i++ {
			field := fields.Field(i)
			value := values.Field(i)
			//fmt.Print("Type:", field.Type, ", ", field.Name, "=", value, "  kind=", value.Kind(), "\n")
			switch field.Type {
			case reflect.TypeOf(""):
				NewString += value.String() + CSVDelimiter
			case reflect.TypeOf([]string{}):
				var NewArray []string = value.Interface().([]string)
				NewString += strings.Join(NewArray, ArrayDelimiter) + CSVDelimiter
			case reflect.TypeOf(in):
				NewString += strconv.FormatInt(value.Int(), 10) + CSVDelimiter
			case reflect.TypeOf([]int{}):
				var IntString string = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(value.Interface().([]int))), ArrayDelimiter), "[]")
				//var NewArray []string = value.Interface().([]int)
				NewString += IntString + CSVDelimiter
			case reflect.TypeOf(i32):
				NewString += strconv.FormatInt(int64(value.Interface().(int32)), 10) + CSVDelimiter
			case reflect.TypeOf(i64):
				NewString += strconv.FormatInt(value.Interface().(int64), 10) + CSVDelimiter
			case reflect.TypeOf(ui):
				NewString += strconv.FormatUint(value.Uint(), 10) + CSVDelimiter
			case reflect.TypeOf(ui32):
				NewString += strconv.FormatUint(value.Uint(), 10) + CSVDelimiter
			case reflect.TypeOf(ui64):
				NewString += strconv.FormatUint(value.Interface().(uint64), 10) + CSVDelimiter
			case reflect.TypeOf(f32):
				NewString += strconv.FormatFloat(value.Float(), 'e', -1, 32) + CSVDelimiter
			case reflect.TypeOf(f64):
				NewString += strconv.FormatFloat(value.Float(), 'e', -1, 64) + CSVDelimiter
			case reflect.TypeOf(b):
				NewString += strconv.FormatBool(value.Bool()) + CSVDelimiter
			case reflect.TypeOf([]S_NetInterface{}):
				var NetInterfaces []S_NetInterface = value.Interface().([]S_NetInterface)
				var NetInterfacesString string
				for _, NetInterface := range NetInterfaces {
					NetInterfacesString = NetInterfacesString + NetInterface.ToString() + ArrayDelimiter
				}
				NewString += NetInterfacesString + CSVDelimiter
			default:
				dealwithErr(fmt.Errorf("StructurToString: error converting " + field.Name + "(" + field.Type.String() + ")"))
			}
		}

		NewCSVArray = append(NewCSVArray, NewString)
	}

	return NewCSVArray
}

func GetHealthSnapshot() S_HealthSnapshot {
	//Get Health infos from local system and return S_HealthSnapshot

	var NewHealthSnapshot S_HealthSnapshot

	//Date and Time
	CurDateTime := time.Now()

	// memory
	vmStat, err := mem.VirtualMemory()
	dealwithErr(err)

	diskStat, err := disk.Usage("/")
	dealwithErr(err)

	// cpu - get CPU number of cores and speed
	cpuStat, err := cpu.Info()
	dealwithErr(err)

	percentage, err := cpu.Percent(0, true)
	dealwithErr(err)

	// host or machine kernel, uptime, platform Info
	hostStat, err := host.Info()
	dealwithErr(err)

	// get interfaces MAC/hardware address
	interfStat, err := net.Interfaces()
	dealwithErr(err)

	NewHealthSnapshot.Date = CurDateTime.Format("2006-01-02")
	NewHealthSnapshot.Time = CurDateTime.Format("15:04:05")

	NewHealthSnapshot.runtimeOS = runtime.GOOS
	NewHealthSnapshot.OS = hostStat.OS
	NewHealthSnapshot.Platform = hostStat.Platform
	NewHealthSnapshot.HostID = hostStat.HostID

	NewHealthSnapshot.Hostname = hostStat.Hostname
	NewHealthSnapshot.Uptime = strconv.FormatUint(hostStat.Uptime, 10)

	//cpuStat[0] -> cpuStat[x] if more than one CPU. Not yet implemented in this script.
	NewHealthSnapshot.CPUindexNumber = strconv.FormatInt(int64(cpuStat[0].CPU), 10)
	NewHealthSnapshot.CPUVendorID = cpuStat[0].VendorID
	NewHealthSnapshot.CPUFamily = cpuStat[0].Family
	NewHealthSnapshot.CPUModelName = cpuStat[0].ModelName
	NewHealthSnapshot.CPUSpeed = strconv.FormatFloat(cpuStat[0].Mhz, 'f', 2, 64) + " MHz"
	NewHealthSnapshot.CPUNumberCores = strconv.FormatInt(int64(cpuStat[0].Cores), 10)

	NewHealthSnapshot.NumberofProcessesRunning = strconv.FormatUint(hostStat.Procs, 10)

	for _, cpupercent := range percentage {
		NewHealthSnapshot.CurrentCPUutilization = append(NewHealthSnapshot.CurrentCPUutilization, strconv.FormatFloat(cpupercent, 'f', 2, 64)+"%")
	}

	NewHealthSnapshot.MemoryTotal = strconv.FormatUint(vmStat.Total, 10)
	NewHealthSnapshot.MemoryFree = strconv.FormatUint(vmStat.Free, 10)
	NewHealthSnapshot.MemoryPercentageUsed = strconv.FormatFloat(vmStat.UsedPercent, 'f', 2, 64)

	NewHealthSnapshot.DiskSpaceTotal = strconv.FormatUint(diskStat.Total, 10)
	NewHealthSnapshot.DiskSpaceUsed = strconv.FormatUint(diskStat.Used, 10)
	NewHealthSnapshot.DiskSpaceFree = strconv.FormatUint(diskStat.Free, 10)
	NewHealthSnapshot.DiskSpaceUsagePercentage = strconv.FormatFloat(diskStat.UsedPercent, 'f', 2, 64) + "%"

	for _, interf := range interfStat {
		var NewNetInterface S_NetInterface //Create new interface

		NewNetInterface.InterfaceName = interf.Name //Interface name
		if interf.HardwareAddr != "" {
			NewNetInterface.MACAddress = interf.HardwareAddr //Interface MAC
		}
		NewNetInterface.Interfacebehaviororflags = append(NewNetInterface.Interfacebehaviororflags, interf.Flags...) //add all flags of interface

		for _, addr := range interf.Addrs {
			NewNetInterface.IPv6_IPv4addresses = append(NewNetInterface.IPv6_IPv4addresses, addr.Addr) //add all IPs of interface
		}

		NewHealthSnapshot.NetInterfaces = append(NewHealthSnapshot.NetInterfaces, NewNetInterface) //Add new interface to NetInterfaces list
	}

	return NewHealthSnapshot
}

func DisplayHealthSnapshot(HealthSnapshot S_HealthSnapshot) {
	//print Health Snapshot to screen

	fmt.Println("Date : " + HealthSnapshot.Date)
	fmt.Println("Time : " + HealthSnapshot.Time)

	fmt.Println("Hostname: " + HealthSnapshot.Hostname)
	fmt.Println("Uptime: " + HealthSnapshot.Uptime)

	fmt.Println("runtimeOS : " + HealthSnapshot.runtimeOS)
	fmt.Println("OS: " + HealthSnapshot.OS)
	fmt.Println("Platform: " + HealthSnapshot.Platform)
	fmt.Println("Host ID(uuid): " + HealthSnapshot.HostID)

	fmt.Println("CPU index number: " + HealthSnapshot.CPUindexNumber)
	fmt.Println("CPU VendorID: " + HealthSnapshot.CPUVendorID)
	fmt.Println("CPU Family: " + HealthSnapshot.CPUFamily)
	fmt.Println("CPU Model Name: " + HealthSnapshot.CPUModelName)
	fmt.Println("CPU Speed: " + HealthSnapshot.CPUSpeed)
	fmt.Println("CPU Number of cores: " + HealthSnapshot.CPUNumberCores)

	fmt.Println("Number of processes running: " + HealthSnapshot.NumberofProcessesRunning)

	for _, cpupercent := range HealthSnapshot.CurrentCPUutilization {
		fmt.Println("Current CPU utilization: " + cpupercent)
	}

	fmt.Println("Total memory: " + HealthSnapshot.MemoryTotal)
	fmt.Println("Free memory: " + HealthSnapshot.MemoryFree)
	fmt.Println("Percentage used memory: " + HealthSnapshot.MemoryPercentageUsed)

	fmt.Println("Total disk space: " + HealthSnapshot.DiskSpaceTotal)
	fmt.Println("Used disk space: " + HealthSnapshot.DiskSpaceUsed)
	fmt.Println("Free disk space: " + HealthSnapshot.DiskSpaceFree)
	fmt.Println("Percentage disk space usage: " + HealthSnapshot.DiskSpaceUsagePercentage)

	for _, interf := range HealthSnapshot.NetInterfaces {
		fmt.Println("Interface : " + interf.InterfaceName)
		fmt.Println("Hardware(MAC) Address: " + interf.MACAddress)
		fmt.Println("IPv6 or IPv4 addresses: " + strings.Join(interf.IPv6_IPv4addresses, ", "))
		fmt.Println("Interface behavior or flags: " + strings.Join(interf.Interfacebehaviororflags, ", "))
	}
}

func (HealthSnapshot *S_HealthSnapshot) HealthSnapshotToDisplayString() string {
	//print Health Snapshot to screen

	var DisplayString string

	DisplayString += "Date : " + HealthSnapshot.Date + "\n"
	DisplayString += "Time : " + HealthSnapshot.Time + "\n"

	DisplayString += "Hostname: " + HealthSnapshot.Hostname + "\n"
	DisplayString += "Uptime: " + HealthSnapshot.Uptime + "\n"

	DisplayString += "runtimeOS : " + HealthSnapshot.runtimeOS + "\n"
	DisplayString += "OS: " + HealthSnapshot.OS + "\n"
	DisplayString += "Platform: " + HealthSnapshot.Platform + "\n"
	DisplayString += "Host ID(uuid): " + HealthSnapshot.HostID + "\n"

	DisplayString += "CPU index number: " + HealthSnapshot.CPUindexNumber + "\n"
	DisplayString += "CPU VendorID: " + HealthSnapshot.CPUVendorID + "\n"
	DisplayString += "CPU Family: " + HealthSnapshot.CPUFamily + "\n"
	DisplayString += "CPU Model Name: " + HealthSnapshot.CPUModelName + "\n"
	DisplayString += "CPU Speed: " + HealthSnapshot.CPUSpeed + "\n"
	DisplayString += "CPU Number of cores: " + HealthSnapshot.CPUNumberCores + "\n"

	DisplayString += "Number of processes running: " + HealthSnapshot.NumberofProcessesRunning + "\n"

	for _, cpupercent := range HealthSnapshot.CurrentCPUutilization {
		DisplayString += "Current CPU utilization: " + cpupercent + "\n"
	}

	DisplayString += "Total memory: " + HealthSnapshot.MemoryTotal + "\n"
	DisplayString += "Free memory: " + HealthSnapshot.MemoryFree + "\n"
	DisplayString += "Percentage used memory: " + HealthSnapshot.MemoryPercentageUsed + "\n"

	DisplayString += "Total disk space: " + HealthSnapshot.DiskSpaceTotal + "\n"
	DisplayString += "Used disk space: " + HealthSnapshot.DiskSpaceUsed + "\n"
	DisplayString += "Free disk space: " + HealthSnapshot.DiskSpaceFree + "\n"
	DisplayString += "Percentage disk space usage: " + HealthSnapshot.DiskSpaceUsagePercentage + "\n"

	for _, interf := range HealthSnapshot.NetInterfaces {
		DisplayString += "Interface : " + interf.InterfaceName + "\n"
		DisplayString += "Hardware(MAC) Address: " + interf.MACAddress + "\n"
		DisplayString += "IPv6 or IPv4 addresses: " + strings.Join(interf.IPv6_IPv4addresses, ", ") + "\n"
		DisplayString += "Interface behavior or flags: " + strings.Join(interf.Interfacebehaviororflags, ", ") + "\n"
	}

	return DisplayString
}

func GetDataPath() string {
	return filepath.Dir((os.Args[0])) + string(os.PathSeparator)
}
func GetPreviousSnapshotFilePath() string {
	PTime := time.Now().AddDate(0, -1, 0)
	month := PTime.Format("01")
	year := PTime.Format("2006")

	basefilename := filepath.Base(os.Args[0])
	basefilename = strings.TrimSuffix(filepath.Base(basefilename), filepath.Ext(basefilename))
	basefilename = basefilename + "_" + year + "-" + month + ".log"

	Snapshotfilepath := GetDataPath() + basefilename

	return Snapshotfilepath
}
func GetCurrentSnapshotFilePath() string {
	month := time.Now().Format("01")
	year := time.Now().Format("2006")

	basefilename := filepath.Base(os.Args[0])
	basefilename = strings.TrimSuffix(filepath.Base(basefilename), filepath.Ext(basefilename))
	basefilename = basefilename + "_" + year + "-" + month + ".log"

	Snapshotfilepath := GetDataPath() + basefilename

	return Snapshotfilepath
}

func AppendStringLineToFile(StringLine, FilePath string) {

	FileObject, err := os.OpenFile(FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		dealwithErr(err)
		return
	}
	defer FileObject.Close()

	if _, err := FileObject.WriteString(StringLine + "\n"); err != nil {
		dealwithErr(err)
	}
	FileObject.Close()
}

func ReadFileToStringArray(FilePath string) []string {

	//Read the whole content of file and return it as string array
	FileObject, err := os.Open(FilePath)
	if err != nil {
		dealwithErr(err)
		return nil
	}
	defer FileObject.Close()

	fileScanner := bufio.NewScanner(FileObject)
	fileScanner.Split(bufio.ScanLines)

	var FileLines []string

	for fileScanner.Scan() {
		FileLines = append(FileLines, fileScanner.Text())
	}
	if err = FileObject.Close(); err != nil {
		dealwithErr(err)
	}

	return FileLines
}

func ReadFileToString(FilePath string) string {
	//Read the whole content of file and return it as string
	FileObject, err := os.ReadFile(FilePath)
	if err != nil {
		dealwithErr(err)
		return ""
	}
	return string(FileObject)
}

func dealwithErr(err error) {
	if err != nil {
		fmt.Println(err)
		//panic(err)
		//os.Exit(-1)
	}
}

func mail(subject string, message string, from string, to []string, cc []string, bcc []string, replyto string) error {

	//Example to send mail
	//err := mail("Subject test "+time.Now().Format("2006-01-02 150405"), "Message test 1\r\nblablablabla\n\rsdfgsdfgsdfgsdfgsdf", "from@example.com", []string{"to1@gmx.fr", "to2@gmx.fr", "to3@gmx.fr", "to4@gmx.fr", "to5@gmx.fr"}, []string{"cc1@test.fr", "cc2@test.fr", "cc3@test.fr"}, []string{"cci1@test.fr", "cci2@test.fr", "cci3@test.fr"}, "reply1@gmx.fr")
	//dealwithErr(err)

	var FullMessage string
	var FullTo []string = to

	//From
	if from == "" {
		return fmt.Errorf("from adress not specified")
	}
	if !strings.Contains(from, "@") && !strings.Contains(from, ".") {
		return fmt.Errorf("invalid from adress")
	}
	FullMessage += "From: " + from + "\r\n"

	//To
	if len(to) < 1 {
		return fmt.Errorf("to adress not specified")
	}
	FullMessage += "To: " + strings.Join(to, "\r\n    ") + "\r\n"

	//Reply-To
	if replyto != "" {
		FullMessage += "Reply-To: " + replyto + "\r\n"
	}

	//Bcc
	if len(bcc) > 0 {
		FullMessage += "Bcc: " + strings.Join(bcc, "\r\n     ") + "\r\n"
		FullTo = append(FullTo, bcc...)
	}

	//Cc
	if len(cc) > 0 {
		FullMessage += "Cc: " + strings.Join(cc, "\r\n    ") + "\r\n"
		FullTo = append(FullTo, cc...)
	}

	//Subject
	FullMessage += "Subject: " + subject + "\r\n\r\n"

	//Message
	FullMessage += message + "\r\n"

	//Attachment
	/*
		AttachmentFilePath := "/your/attachment-filepath"

		AttachmentDelimeter := "**=myohmy689407924327"
		AttachmentFilename := "logFileUI.txt"

		AttachmentInfos := "\r\n--" + AttachmentDelimeter + "\r\n"
		AttachmentInfos += "MIME-Version: 1.0\r\n"
		AttachmentInfos += "Content-Type: text/plain; charset=\"utf-8\"\r\n"
		AttachmentInfos += "Content-Transfer-Encoding: base64\r\n"
		AttachmentInfos += "Content-Disposition: attachment;filename=\"" + AttachmentFilename + "\"\r\n"

		rawFile, fileErr := ioutil.ReadFile(AttachmentFilePath)
		if fileErr != nil {
			dealwithErr(fileErr)
		}

		AttachmentFileContent := base64.StdEncoding.EncodeToString(rawFile)
		FullMessage += AttachmentInfos + AttachmentFileContent
	*/

	// smtp server configuration.
	smtpHost := "smtp.mailtrap.io"
	smtpPort := "2525"

	smtplogin := "ccd2ae332055ea"
	smtppassword := "0fc1b7ccc90d2e"

	// Authentication. // func PlainAuth(identity, username, password, host string) Auth
	auth := smtp.PlainAuth("", smtplogin, smtppassword, smtpHost)

	// Sending email. // func SendMail(addr string, a Auth, from string, to []string, msg []byte) error
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, FullTo, []byte(FullMessage))
	//err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, FullTo, []byte(FullMessage))
	if err != nil {
		return err
		/*} else {
		fmt.Println("1: ", smtpHost+":"+smtpPort)
		fmt.Println("2: ", auth)
		fmt.Println("3: ", from)
		fmt.Println("4: ", strings.Join(FullTo, "\r\n"))
		fmt.Println("5: ", strings.Join(cc, "\r\n"))
		fmt.Println("6: ", strings.Join(bcc, "\r\n"))
		fmt.Println("7: ", NewLineChar+FullMessage)
		*/
	}

	//no error
	return nil
}

func LiveView() {
	list := tview.NewList()

	// Ajoutez des éléments à la liste
	list.AddItem("Element 1", "", '1', nil)
	list.AddItem("Element 2", "", '2', nil)
	list.AddItem("Element 3", "", '3', nil)

	// Configurez l'apparence de la liste
	list.SetMainTextColor(tcell.ColorWhite)
	list.SetSelectedTextColor(tcell.ColorYellow)
	list.SetSelectedBackgroundColor(tcell.ColorBlue)

	// Gérez les événements de sélection de l'élément
	list.SetSelectedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		fmt.Println("L'élément sélectionné est", mainText)
	})

	// Ajoutez la liste à la vue principale de l'application
	app.SetRoot(list, true)

}

func LiveViewtest2() {

	var MyApp = tview.NewApplication()

	MyApp.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 113 {
			MyApp.Stop()
		}
		return event
	})

	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}
	menu := newPrimitive("Menu")
	main := newPrimitive("Main content")
	sideBar := newPrimitive("Side Bar")

	grid := tview.NewGrid().
		SetRows(3, 0, 3).
		SetColumns(40, 0, 30).
		SetBorders(true).
		AddItem(newPrimitive("Header"), 0, 0, 1, 3, 0, 0, false).
		AddItem(newPrimitive("Footer"), 2, 0, 1, 3, 0, 0, false)

	// Layout for screens narrower than 100 cells (menu and side bar are hidden).
	grid.AddItem(menu, 0, 0, 0, 0, 0, 0, false).
		AddItem(main, 1, 0, 1, 3, 0, 0, false).
		AddItem(sideBar, 0, 0, 0, 0, 0, 0, false)

	// Layout for screens wider than 100 cells.
	grid.AddItem(menu, 1, 0, 1, 1, 0, 100, false).
		AddItem(main, 1, 1, 1, 1, 0, 100, false).
		AddItem(sideBar, 1, 2, 1, 1, 0, 100, false)

	if err := tview.NewApplication().SetRoot(grid, true).SetFocus(grid).Run(); err != nil {
		panic(err)
	}

	/*
		var MyApp = tview.NewApplication()
		var text = tview.NewTextView().
			SetTextColor(tcell.ColorGreen).
			SetText("(q) to quit")

		if err := MyApp.SetRoot(text, true).EnableMouse(true).Run(); err != nil {
			panic(err)
		}

		MyApp.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Rune() == 113 {
				MyApp.Stop()
			}
			return event
		})
	*/
	/*
		MyApp := tview.NewApplication()

		if err := MyApp.SetRoot(tview.NewBox(), true).EnableMouse(true).Run(); err != nil {
			panic(err)
		}
	*/

	/*
		flex := tview.NewFlex().
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Left (1/2 x width of Top)"), 0, 1, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(tview.NewBox().SetBorder(true).SetTitle("Top"), 0, 1, false).
				AddItem(tview.NewBox().SetBorder(true).SetTitle("Middle (3 x height of Top)"), 0, 3, false).
				AddItem(tview.NewBox().SetBorder(true).SetTitle("Bottom (5 rows)"), 5, 1, false), 0, 2, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Right (20 cols)"), 20, 1, false)
		if err := MyApp.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
			panic(err)
		}
	*/
}

func FilesView() {
	BaseDir := GetDataPath()

	f, err := os.Open(BaseDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	files, err := f.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range files {
		fmt.Println(v.Name(), v.IsDir())
	}

}

func StartMonitor() {
	//Run in background to GetHealthSnapshot, send alerts and reports
	for {
		HealthSnapshot := GetHealthSnapshot()
		Snapshotfilepath := GetCurrentSnapshotFilePath()
		AppendStringLineToFile(HealthSnapshot.ToLogLine(), Snapshotfilepath)
		duration := time.Second
		time.Sleep(duration * 3600)
	}
}

func main() {

	// the first argument i.e. program name is excluded
	if len(os.Args[1:]) > 0 {
		for _, Argument := range os.Args[1:] {
			//Argument = strings.ReplaceAll(Argument,"--","-")
			switch strings.ToLower(Argument) {
			case "--help", "-help", "/?", "-h":
				fmt.Println("Record computer health and send mail reports.")
				fmt.Println("\nUsage: " + filepath.Base(os.Args[0]) + " (<option>)")
				fmt.Println("Options:")
				fmt.Println("-help, -h or /? : Display this help.")
				fmt.Println("-service, -s, or /s : start in background and send mail reports")
				fmt.Println("-live, -l or /l : live view of computer health")
				fmt.Println("-open, -o or /o : view saved log of computer health")
			case "--service", "-service", "-s", "/s":
				StartMonitor()
			case "--live", "-live", "-l", "/l", "--gui", "-gui", "/gui":
				LiveView()
			case "--open", "-open", "-o", "/o", "/open":
				FilesView()
			default:
				LiveView()
			}
		}
	} else {
		//var HealthSnapshot S_HealthSnapshot = GetHealthSnapshot()
		//fmt.Print("test: " + HealthSnapshot.GetField("OS"))

		//fmt.Print("HS - " + HealthSnapshot.ToLogLine())

		//fmt.Println("GetCurrentSnapshotFilePath: " + GetCurrentSnapshotFilePath())
		//fmt.Println("GetPreviousSnapshotFilePath: " + GetPreviousSnapshotFilePath())

		//fmt.Print(HealthSnapshot.HealthSnapshotToDisplayString())

		FilesView()

		fmt.Print("Press 'Enter' to continue...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}
