package routes

import (
	"encoding/json"
	"localdev/main/config"
	"localdev/main/services"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/grpc"

	pb "localdev/main/proto/user_matching"
)

func GetMatches(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	userService := services.Locator.UserService()

	userID := userService.GetCurrentUserID(ctx)

	if userID == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	conn, err := grpc.Dial(config.MatcherAddress(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	defer conn.Close()

	matcherClient := pb.NewMatcherClient(conn)

	resp, err := matcherClient.GetMatches(ctx, &pb.GetMatchesRequest{UserId: userID})
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	out, err := json.Marshal(resp)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(out))
}
