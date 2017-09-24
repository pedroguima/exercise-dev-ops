package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"

	mdb "github.com/msackman/gomdb"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/golang/protobuf/proto"

	pb "github.com/tokencard/dev-ops-exercise/src/routeguide"
)

var (
	// lmdb
	env *mdb.Env
	dbi mdb.DBI

	backendPort = flag.Int("backend", 9999, "DB backend address")
)

type backendServer struct {
	savedFeatures []*pb.Feature
	routeNotes    map[string][]*pb.RouteNote
	backend       pb.RouteGuideClient
}
var _ = (*pb.RouteGuideServer)(nil)

type SavedFeature struct {
	ID       string // a date including nanoseconds ("2006-01-02T15:04:05.999999999")
	Name     string
	Location *pb.Point
}

// GetFeature returns the feature at the given point.
func (bs *backendServer) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
	for _, feature := range bs.savedFeatures {
		if proto.Equal(feature.Location, point) {
			return feature, nil
		}
	}
	// No feature was found, return an unnamed feature
	return &pb.Feature{Location: point}, nil
}

func (bs *backendServer) SaveFeature(ctx context.Context, feature *pb.Feature) (*pb.SavedResult, error) {
	var savedResult *pb.SavedResult

	if err := putInMemDB(feature); err != nil {
		log.Printf("putInMemDB() error: %s\n", err)
		savedResult = &pb.SavedResult{
			Ok: false,
		}
		return savedResult, err
	}
	savedResult = &pb.SavedResult{
		Ok: true,
	}

	savedFeatures, err := getAllSavedFeatures()
	if err != nil {
		log.Printf("getAllSavedFeatures() error: %s\n", err)
	}

	for _, savedFeature := range savedFeatures {
		fmt.Printf("SAVED FEATURE\t\tFeature[%v] has Name: %v\n and Location: %v\n", savedFeature.ID, savedFeature.Name, savedFeature.Location)
	}

	return savedResult, nil
}

// ListFeatures lists all features contained within the given bounding Rectangle.
func (bs *backendServer) ListFeatures(rect *pb.Rectangle, stream pb.RouteGuide_ListFeaturesServer) error {
	// DO NOTHING
	return nil
}

// RecordRoute records a route composited of a sequence of points.
//
// It gets a stream of points, and responds with statistics about the "trip":
// number of points,  number of known features visited, total distance traveled, and
// total time spent.
func (bs *backendServer) RecordRoute(stream pb.RouteGuide_RecordRouteServer) error {
	// DO NOTHING
	return nil
}

// RouteChat receives a stream of message/location pairs, and responds with a stream of all
// previous messages at each of those locations.
func (bs *backendServer) RouteChat(stream pb.RouteGuide_RouteChatServer) error {
	// DO NOTHING
	return nil
}

// loadFeatures loads features from a JSON file.
func (bs *backendServer) loadFeatures(filePath string) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to load default features: %v", err)
	}
	if err := json.Unmarshal(file, &bs.savedFeatures); err != nil {
		log.Fatalf("Failed to load default features: %v", err)
	}
}

func main() {
	flag.Parse()

	// Setup of the LMDB DB on disk
	path := "./mdb_data"
	if memDBFileInfo, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Printf("os.Mkdir() error: %s\n", err)
		}
	} else {
		log.Printf("os.Stat() path has size: %d\n", memDBFileInfo.Size())
	}

	env, _ = mdb.NewEnv()
	env.SetMapSize(1 << 30) // max file size: 1 GB
	env.Open(path, 0, 0664)
	defer env.Close()

	txn, _ := env.BeginTxn(nil, 0) //it is a root transaction, parent is nil
	dbi, _ = txn.DBIOpen(nil, 0)   //'nil' here indicates a single db will be used
	defer env.DBIClose(dbi)
	txn.Commit()
	stat, _ := env.Stat()
	log.Printf("The DB has %d entries\n", stat.Entries)
	///////////// LMDB setup

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *backendPort)) // RPC port
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	g := grpc.NewServer()
	pb.RegisterRouteGuideServer(g, new(backendServer))
	log.Printf("backend listening on  :%d...\n", *backendPort)
	g.Serve(lis)
}

// LMDB utilities
func putInMemDB(feature *pb.Feature) error {
		savedFeature := &SavedFeature{
			ID:       time.Now().Format("2006-01-02T15:04:05.999999999"),
			Name:     feature.Name,
			Location: feature.Location,
		}
		log.Printf("savedFeature is: %#v\n", savedFeature)

		// Gobbed (a bytes.Buffer) is an io.Reader and an io.Writer
		Gobbed, err := EncodeGob(savedFeature)
		if err != nil {
			log.Printf("put EncodeGob() error: %s\n", err)
		}

		txn, err := env.BeginTxn(nil, 0)
		if err != nil {
			log.Printf("put BeginTxn() error: %s\n", err)
		}
		key := []byte(fmt.Sprintf("%s", savedFeature.ID))
		val := Gobbed.Bytes()
		if err := txn.Put(dbi, key, val, 0); err != nil {
			log.Printf("txn.Put() error: %s\n", err)
		}
		txn.Commit()

	return nil
}

func getAllSavedFeatures() ([]SavedFeature, error) {
	savedFeature := &SavedFeature{}
	savedFeatures := []SavedFeature{}

	// scan the database
	stat, _ := env.Stat()
	fmt.Printf("The DB has %d entries\n", stat.Entries)
	txn, _ := env.BeginTxn(nil, mdb.RDONLY)
	defer txn.Abort()

	cursor, _ := txn.CursorOpen(dbi)
	defer cursor.Close()

	entries := stat.Entries
	for i := 0; i < int(entries); i++ {
		key, val, err := cursor.Get(nil, nil, mdb.NEXT_NODUP)
		if err == mdb.NotFound {
			break
		}
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		if err = DecodeGob(savedFeature, val); err != nil {
			log.Println(err)
		}
		savedFeature.ID = string(key)
		//savedEventAsTime, _ := time.Parse("2006-01-02T15:04:05.999999999", savedEvent.ID)
		savedFeatures = append(savedFeatures, *savedFeature)
	}
	return savedFeatures, nil
}

func EncodeGob(src interface{}) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(src); err != nil {
		return nil, err
	}
	//return buf.Bytes(), nil
	return buf, nil
}

func DecodeGob(dst interface{}, src []byte) error {
	dec := gob.NewDecoder(bytes.NewBuffer(src))
	if err := dec.Decode(dst); err != nil {
		return err
	}
	return nil
}
