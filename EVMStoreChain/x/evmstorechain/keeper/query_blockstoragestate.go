package keeper

import (
	"context"

	"EVMStoreChain/x/evmstorechain/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) BlockstoragestateAll(goCtx context.Context, req *types.QueryAllBlockstoragestateRequest) (*types.QueryAllBlockstoragestateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var blockstoragestates []types.Blockstoragestate
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	blockstoragestateStore := prefix.NewStore(store, types.KeyPrefix(types.BlockstoragestateKeyPrefix))

	pageRes, err := query.Paginate(blockstoragestateStore, req.Pagination, func(key []byte, value []byte) error {
		var blockstoragestate types.Blockstoragestate
		if err := k.cdc.Unmarshal(value, &blockstoragestate); err != nil {
			return err
		}

		blockstoragestates = append(blockstoragestates, blockstoragestate)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllBlockstoragestateResponse{Blockstoragestate: blockstoragestates, Pagination: pageRes}, nil
}

func (k Keeper) Blockstoragestate(goCtx context.Context, req *types.QueryGetBlockstoragestateRequest) (*types.QueryGetBlockstoragestateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetBlockstoragestate(
		ctx,
		req.Blocknumber,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetBlockstoragestateResponse{Blockstoragestate: val}, nil
}
