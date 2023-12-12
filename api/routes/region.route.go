package routes

import (
	"context"

	"github.com/AhmetSBulbul/quarterback-server/api/pb/commonpb"
	"github.com/AhmetSBulbul/quarterback-server/api/pb/regionpb"
)

type RegionService struct{}

func (r *RegionService) ListCountry(_ context.Context, _ *commonpb.Empty) (*regionpb.CountryListResponse, error) {
	panic("not implemented") // TODO: Implement
}

// Should we use query and pagination?
func (r *RegionService) ListCity(_ context.Context, _ *commonpb.GetByIdRequest) (*regionpb.CityListResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (r *RegionService) ListDistrict(_ context.Context, _ *commonpb.GetByIdRequest) (*regionpb.DistrictListResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (r *RegionService) mustEmbedUnimplementedRegionServiceServer() {
	panic("not implemented") // TODO: Implement
}
