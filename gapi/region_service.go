package gapi

import (
	"context"
	"database/sql"
	"log"

	"github.com/AhmetSBulbul/quarterback-server/pb/commonpb"
	"github.com/AhmetSBulbul/quarterback-server/pb/regionpb"
	"google.golang.org/grpc/codes"
)

type RegionService struct {
	db *sql.DB
	regionpb.UnimplementedRegionServiceServer
}

func (r *RegionService) ListCountry(ctx context.Context, _ *commonpb.Empty) (*regionpb.CountryListResponse, error) {
	query := "SELECT id, name FROM country"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, gerr(codes.Internal, err)
	}

	defer rows.Close()

	countries := []*regionpb.Country{}

	for rows.Next() {
		var country regionpb.Country
		err := rows.Scan(&country.Id, &country.Name)

		if err != nil {
			log.Println(err.Error())
			return nil, gerr(codes.Internal, err)
		}

		countries = append(countries, &country)
	}

	return &regionpb.CountryListResponse{
		Countries: countries,
	}, nil
}

// Should we use query and pagination?
func (r *RegionService) ListCity(ctx context.Context, in *commonpb.GetByIdRequest) (*regionpb.CityListResponse, error) {
	query := "SELECT id, name FROM city WHERE countryID = ?"
	rows, err := r.db.QueryContext(ctx, query, in.Id)
	if err != nil {
		return nil, gerr(codes.Internal, err)
	}

	defer rows.Close()

	cities := []*regionpb.City{}

	for rows.Next() {
		var city regionpb.City
		err := rows.Scan(&city.Id, &city.Name)

		if err != nil {
			log.Println(err.Error())
			return nil, gerr(codes.Internal, err)
		}

		cities = append(cities, &city)
	}

	return &regionpb.CityListResponse{
		Cities: cities,
	}, nil
}

func (r *RegionService) ListDistrict(ctx context.Context, in *commonpb.GetByIdRequest) (*regionpb.DistrictListResponse, error) {
	query := "SELECT id, name FROM district WHERE cityID = ?"
	rows, err := r.db.QueryContext(ctx, query, in.Id)
	if err != nil {
		return nil, gerr(codes.Internal, err)
	}

	defer rows.Close()

	districts := []*regionpb.District{}

	for rows.Next() {
		var district regionpb.District
		err := rows.Scan(&district.Id, &district.Name)

		if err != nil {
			log.Println(err.Error())
			return nil, gerr(codes.Internal, err)
		}

		districts = append(districts, &district)
	}

	return &regionpb.DistrictListResponse{
		Districts: districts,
	}, nil
}

func (r *RegionService) GetRegionByDistrictId(ctx context.Context, in *commonpb.GetByIdRequest) (*regionpb.Region, error) {
	var country regionpb.Country
	var city regionpb.City
	var district regionpb.District

	query := `SELECT
		co.id,
		co.name,
		c.id,
		c.name,
		c.countryID,
		d.id,
		d.name,
		d.cityID
		FROM district d
		INNER JOIN city c ON c.id = d.cityID
		INNER JOIN country co ON co.id = c.countryID
		WHERE d.id = ?`
	row := r.db.QueryRowContext(ctx, query, in.Id)

	err := row.Scan(&country.Id, &country.Name, &city.Id, &city.Name, &city.CountryId, &district.Id, &district.Name, &district.CityId)
	if err != nil {
		return nil, gerr(codes.Internal, err)
	}

	var region regionpb.Region

	region.Country = &country
	region.City = &city
	region.District = &district

	return &region, nil
}
