package query

import (
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
)

// ibc_ParamsRPC returns the distribution params
func ibc_ClientParamsRPC(q *Query) (*clienttypes.QueryClientParamsResponse, error) {
	req := &clienttypes.QueryClientParamsRequest{}
	queryClient := clienttypes.NewQueryClient(q.Client)
	ctx, cancel := q.GetQueryContext()
	defer cancel()
	res, err := queryClient.ClientParams(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ibc_ClientStateRPC returns the state of the specified IBC client.
func ibc_ClientStateRPC(q *Query, clientId string) (*clienttypes.QueryClientStateResponse, error) {
	req := &clienttypes.QueryClientStateRequest{ClientId: clientId}
	queryClient := clienttypes.NewQueryClient(q.Client)
	ctx, cancel := q.GetQueryContext()
	defer cancel()
	res, err := queryClient.ClientState(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
