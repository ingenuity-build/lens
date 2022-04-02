package staking

import (
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/strangelove-ventures/lens/client/query"
)

// ParamsRPC returns the staking params
func ParamsRPC(q *query.Query) (*stakingTypes.QueryParamsResponse, error) {
	req := &stakingTypes.QueryParamsRequest{}
	queryClient := stakingTypes.NewQueryClient(q.Client)
	ctx, cancel := q.GetQueryContext()
	defer cancel()
	res, err := queryClient.Params(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ParamsRPC returns the distribution params
func PoolRPC(q *query.Query) (*stakingTypes.QueryPoolResponse, error) {
	req := &stakingTypes.QueryPoolRequest{}
	queryClient := stakingTypes.NewQueryClient(q.Client)
	ctx, cancel := q.GetQueryContext()
	defer cancel()
	res, err := queryClient.Pool(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// DelegationRPC returns the delegations to a particular validator
func DelegationRPC(q *query.Query, delegator, validator string) (*stakingTypes.QueryDelegationResponse, error) {
	// ensure the delegator parameter is a valid account address
	_, err := q.Client.DecodeBech32AccAddr(delegator)
	if err != nil {
		return nil, err
	}
	// ensure the validator parameter is a valid validator address
	_, err = q.Client.DecodeBech32ValAddr(validator)
	if err != nil {
		return nil, err
	}
	queryClient := stakingTypes.NewQueryClient(q.Client)
	req := &stakingTypes.QueryDelegationRequest{
		DelegatorAddr: delegator,
		ValidatorAddr: validator,
	}
	ctx, cancel := q.GetQueryContext()
	defer cancel()
	res, err := queryClient.Delegation(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// DelegatorDelegationsRPC returns all the delegations for a given delegator
func DelegatorDelegationsRPC(q *query.Query, delegator string) (*stakingTypes.QueryDelegatorDelegationsResponse, error) {
	// ensure the delegator parameter is a valid account address
	_, err := q.Client.DecodeBech32AccAddr(delegator)
	if err != nil {
		return nil, err
	}
	queryClient := stakingTypes.NewQueryClient(q.Client)
	req := &stakingTypes.QueryDelegatorDelegationsRequest{
		DelegatorAddr: delegator,
		Pagination:    q.Options.Pagination,
	}
	ctx, cancel := q.GetQueryContext()
	defer cancel()
	res, err := queryClient.DelegatorDelegations(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// DelegationsRPC returns all the delegations
func DelegatorUnbondingDelegationsRPC(q *query.Query, delegator string) (*stakingTypes.QueryDelegatorUnbondingDelegationsResponse, error) {
	// ensure the delegator parameter is a valid account address
	_, err := q.Client.DecodeBech32AccAddr(delegator)
	if err != nil {
		return nil, err
	}
	queryClient := stakingTypes.NewQueryClient(q.Client)
	req := &stakingTypes.QueryDelegatorUnbondingDelegationsRequest{
		DelegatorAddr: delegator,
		Pagination:    q.Options.Pagination,
	}
	ctx, cancel := q.GetQueryContext()
	defer cancel()
	res, err := queryClient.DelegatorUnbondingDelegations(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ValidatorsRPC returns all the validators for a given status
func ValidatorsRPC(q *query.Query, status string) (*stakingTypes.QueryValidatorsResponse, error) {
	queryClient := stakingTypes.NewQueryClient(q.Client)
	req := &stakingTypes.QueryValidatorsRequest{
		Status:     status,
		Pagination: q.Options.Pagination,
	}
	ctx, cancel := q.GetQueryContext()
	defer cancel()
	res, err := queryClient.Validators(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ValidatorRPC returns all the validator for a given address
func ValidatorRPC(q *query.Query, address string) (*stakingTypes.QueryValidatorResponse, error) {
	// ensure the validator parameter is a valid validator address
	_, err := q.Client.DecodeBech32ValAddr(address)
	if err != nil {
		return nil, err
	}
	queryClient := stakingTypes.NewQueryClient(q.Client)
	req := &stakingTypes.QueryValidatorRequest{
		ValidatorAddr: address,
	}
	ctx, cancel := q.GetQueryContext()
	defer cancel()
	res, err := queryClient.Validator(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ValidatorDelegationsRPC returns all the delegations for a validator
func ValidatorDelegationsRPC(q *query.Query, validator string) (*stakingTypes.QueryValidatorDelegationsResponse, error) {
	// ensure the validator parameter is a valid validator address
	_, err := q.Client.DecodeBech32ValAddr(validator)
	if err != nil {
		return nil, err
	}
	queryClient := stakingTypes.NewQueryClient(q.Client)
	req := &stakingTypes.QueryValidatorDelegationsRequest{
		ValidatorAddr: validator,
		Pagination:    q.Options.Pagination,
	}
	ctx, cancel := q.GetQueryContext()
	defer cancel()
	res, err := queryClient.ValidatorDelegations(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ValidatorUnbondingDelegationsRPC returns all the unbonding delegations for a validator
func ValidatorUnbondingDelegationsRPC(q *query.Query, validator string) (*stakingTypes.QueryValidatorUnbondingDelegationsResponse, error) {
	// ensure the validator parameter is a valid validator address
	_, err := q.Client.DecodeBech32ValAddr(validator)
	if err != nil {
		return nil, err
	}
	queryClient := stakingTypes.NewQueryClient(q.Client)
	req := &stakingTypes.QueryValidatorUnbondingDelegationsRequest{
		ValidatorAddr: validator,
		Pagination:    q.Options.Pagination,
	}
	ctx, cancel := q.GetQueryContext()
	defer cancel()
	res, err := queryClient.ValidatorUnbondingDelegations(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ValidatorUnbondingDelegationsRPC returns all the unbonding delegations for a validator
func RelegationsRPC(q *query.Query, delegator string, src_validator string, dst_validator string) (*stakingTypes.QueryRedelegationsResponse, error) {
	// ensure the addresses are valid
	_, err := q.Client.DecodeBech32AccAddr(delegator)
	if err != nil {
		return nil, err
	}
	_, err = q.Client.DecodeBech32ValAddr(src_validator)
	if err != nil {
		return nil, err
	}
	_, err = q.Client.DecodeBech32ValAddr(dst_validator)
	if err != nil {
		return nil, err
	}
	queryClient := stakingTypes.NewQueryClient(q.Client)
	req := &stakingTypes.QueryRedelegationsRequest{
		DelegatorAddr:    delegator,
		SrcValidatorAddr: src_validator,
		DstValidatorAddr: dst_validator,
		Pagination:       q.Options.Pagination,
	}
	ctx, cancel := q.GetQueryContext()
	defer cancel()
	res, err := queryClient.Redelegations(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// DelegatorValidatorsRPC returns all the validators for a given delegator
func DelegatorValidatorsRPC(q *query.Query, delegator string) (*stakingTypes.QueryDelegatorValidatorsResponse, error) {
	// ensure the delegator parameter is a valid account address
	_, err := q.Client.DecodeBech32AccAddr(delegator)
	if err != nil {
		return nil, err
	}
	queryClient := stakingTypes.NewQueryClient(q.Client)
	req := &stakingTypes.QueryDelegatorValidatorsRequest{
		DelegatorAddr: delegator,
		Pagination:    q.Options.Pagination,
	}
	ctx, cancel := q.GetQueryContext()
	defer cancel()
	res, err := queryClient.DelegatorValidators(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// DelegatorValidatorRPC returns the validators for a given delegator
func DelegatorValidatorRPC(q *query.Query, delegator string, validator string) (*stakingTypes.QueryDelegatorValidatorResponse, error) {
	// ensure the delegator parameter is a valid account address
	_, err := q.Client.DecodeBech32AccAddr(delegator)
	if err != nil {
		return nil, err
	}
	_, err = q.Client.DecodeBech32ValAddr(validator)
	if err != nil {
		return nil, err
	}
	queryClient := stakingTypes.NewQueryClient(q.Client)
	req := &stakingTypes.QueryDelegatorValidatorRequest{
		DelegatorAddr: delegator,
		ValidatorAddr: validator,
	}
	ctx, cancel := q.GetQueryContext()
	defer cancel()
	res, err := queryClient.DelegatorValidator(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// DelegatorValidatorRPC returns the validators for a given delegator
func HistoricalInfoRPC(q *query.Query, height int64) (*stakingTypes.QueryHistoricalInfoResponse, error) {
	queryClient := stakingTypes.NewQueryClient(q.Client)
	req := &stakingTypes.QueryHistoricalInfoRequest{
		Height: height,
	}
	ctx, cancel := q.GetQueryContext()
	defer cancel()
	res, err := queryClient.HistoricalInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
