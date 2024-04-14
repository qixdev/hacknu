package main

import (
	"log"
	"net/http"
)

func main() {

	//if err := godotenv.Load(); err != nil {
	//	log.Fatalf("Error loading .env file: %v", err)
	//}
	//
	//mongoURI := os.Getenv("MONGO")
	//if mongoURI == "" {
	//	log.Fatal("MONGO environment variable not set")
	//	return
	//}

	//serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	//opts := options.Client().ApplyURI("mongodb+srv://Zhalgas:ZHASIK7799ufc@cluster1.llliwo7.mongodb.net/?retryWrites=true&w=majority&appName=Cluster1").SetServerAPIOptions(serverAPI)
	//client, err := mongo.Connect(context.TODO(), opts)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//defer func() {
	//	if err = client.Disconnect(context.TODO()); err != nil {
	//		log.Println(err)
	//		return
	//	}
	//}()
	//if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
	//	log.Println(err)
	//	return
	//}
	//fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	mux := http.NewServeMux()
	mux.HandleFunc("/teacher", TeacherPage)
	mux.HandleFunc("/student", StudentPage)
	mux.HandleFunc("/parent", ParentPage)

	if err := http.ListenAndServe(":4000", mux); err != nil {
		log.Println(err)
		return
	}
}
