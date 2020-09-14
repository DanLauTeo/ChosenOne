package routes

import (
	"localdev/main/config"
	pb "localdev/main/proto/user_matching"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/grpc"
)

func CheckCronHeader(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header["X-Appengine-Cron"] != nil {
			handler(w, r)
		} else {
			w.WriteHeader(http.StatusForbidden)
		}
	}
}

var HandleRecalcUserMatches = CheckCronHeader(handleRecalcUserMatches)

func handleRecalcUserMatches(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	conn, err := grpc.Dial(config.MatcherAddress(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	defer conn.Close()

	matcherClient := pb.NewMatcherClient(conn)

	if _, err := matcherClient.RecalcScaNN(ctx, &pb.Empty{}); err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	w.WriteHeader(http.StatusOK)
}
