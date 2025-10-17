package app

import (
    "context"
    . "rpi-test/pkg/openapi/v1"
)

type Server struct {
    app *App
}

var _ StrictServerInterface = (*Server)(nil)

func newServer(app *App) *Server {
    return &Server{
        app: app,
    }
}


// PostApiAstanaLrtBuilding1 operation middleware

// (POST )
func (s *Server) PostApiAstanaLrtBuilding1(ctx context.Context, request PostApiAstanaLrtBuilding1RequestObject) (PostApiAstanaLrtBuilding1ResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiAstanaLrtBuilding1200JSONResponse{}, nil
}

// PostApiAstanaLrtStreetlights1 operation middleware

// (POST )
func (s *Server) PostApiAstanaLrtStreetlights1(ctx context.Context, request PostApiAstanaLrtStreetlights1RequestObject) (PostApiAstanaLrtStreetlights1ResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiAstanaLrtStreetlights1200JSONResponse{}, nil
}

// PostApiAstanaLrtTl1 operation middleware

// (POST )
func (s *Server) PostApiAstanaLrtTl1(ctx context.Context, request PostApiAstanaLrtTl1RequestObject) (PostApiAstanaLrtTl1ResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiAstanaLrtTl1200JSONResponse{}, nil
}

// PostApiBusinessCenterBuilding1 operation middleware

// (POST )
func (s *Server) PostApiBusinessCenterBuilding1(ctx context.Context, request PostApiBusinessCenterBuilding1RequestObject) (PostApiBusinessCenterBuilding1ResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiBusinessCenterBuilding1200JSONResponse{}, nil
}

// PostApiGovernmentComplexBuilding1 operation middleware

// (POST )
func (s *Server) PostApiGovernmentComplexBuilding1(ctx context.Context, request PostApiGovernmentComplexBuilding1RequestObject) (PostApiGovernmentComplexBuilding1ResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiGovernmentComplexBuilding1200JSONResponse{}, nil
}

// PostApiGovernmentComplexBuilding2 operation middleware

// (POST )
func (s *Server) PostApiGovernmentComplexBuilding2(ctx context.Context, request PostApiGovernmentComplexBuilding2RequestObject) (PostApiGovernmentComplexBuilding2ResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiGovernmentComplexBuilding2200JSONResponse{}, nil
}

// PostApiGovernmentComplexBuilding3 operation middleware

// (POST )
func (s *Server) PostApiGovernmentComplexBuilding3(ctx context.Context, request PostApiGovernmentComplexBuilding3RequestObject) (PostApiGovernmentComplexBuilding3ResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiGovernmentComplexBuilding3200JSONResponse{}, nil
}

// PostApiGovernmentComplexBuilding4 operation middleware

// (POST )
func (s *Server) PostApiGovernmentComplexBuilding4(ctx context.Context, request PostApiGovernmentComplexBuilding4RequestObject) (PostApiGovernmentComplexBuilding4ResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiGovernmentComplexBuilding4200JSONResponse{}, nil
}

// PostApiHospitalBarrier operation middleware

// (POST )
func (s *Server) PostApiHospitalBarrier(ctx context.Context, request PostApiHospitalBarrierRequestObject) (PostApiHospitalBarrierResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiHospitalBarrier200JSONResponse{}, nil
}

// PostApiHospitalBuilding1 operation middleware

// (POST )
func (s *Server) PostApiHospitalBuilding1(ctx context.Context, request PostApiHospitalBuilding1RequestObject) (PostApiHospitalBuilding1ResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiHospitalBuilding1200JSONResponse{}, nil
}

// PostApiHospitalFire operation middleware

// (POST )
func (s *Server) PostApiHospitalFire(ctx context.Context, request PostApiHospitalFireRequestObject) (PostApiHospitalFireResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiHospitalFire200JSONResponse{}, nil
}

// PostApiHospitalFireAlarm operation middleware

// (POST )
func (s *Server) PostApiHospitalFireAlarm(ctx context.Context, request PostApiHospitalFireAlarmRequestObject) (PostApiHospitalFireAlarmResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiHospitalFireAlarm200JSONResponse{}, nil
}

// PostApiHospitalFireTruck operation middleware

// (POST )
func (s *Server) PostApiHospitalFireTruck(ctx context.Context, request PostApiHospitalFireTruckRequestObject) (PostApiHospitalFireTruckResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiHospitalFireTruck200JSONResponse{}, nil
}

// PostApiHospitalSiren operation middleware

// (POST )
func (s *Server) PostApiHospitalSiren(ctx context.Context, request PostApiHospitalSirenRequestObject) (PostApiHospitalSirenResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiHospitalSiren200JSONResponse{}, nil
}

// PostApiSportsComplexStadiumLights operation middleware

// (POST )
func (s *Server) PostApiSportsComplexStadiumLights(ctx context.Context, request PostApiSportsComplexStadiumLightsRequestObject) (PostApiSportsComplexStadiumLightsResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiSportsComplexStadiumLights200JSONResponse{}, nil
}

// PostApiSportsComplexTl1 operation middleware

// (POST )
func (s *Server) PostApiSportsComplexTl1(ctx context.Context, request PostApiSportsComplexTl1RequestObject) (PostApiSportsComplexTl1ResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiSportsComplexTl1200JSONResponse{}, nil
}

// PostApiWaterPumpBuilding1 operation middleware

// (POST )
func (s *Server) PostApiWaterPumpBuilding1(ctx context.Context, request PostApiWaterPumpBuilding1RequestObject) (PostApiWaterPumpBuilding1ResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiWaterPumpBuilding1200JSONResponse{}, nil
}

// PostApiWaterPumpPump2 operation middleware

// (POST )
func (s *Server) PostApiWaterPumpPump2(ctx context.Context, request PostApiWaterPumpPump2RequestObject) (PostApiWaterPumpPump2ResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiWaterPumpPump2200JSONResponse{}, nil
}

// PostApiWaterPumpPump3 operation middleware

// (POST )
func (s *Server) PostApiWaterPumpPump3(ctx context.Context, request PostApiWaterPumpPump3RequestObject) (PostApiWaterPumpPump3ResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiWaterPumpPump3200JSONResponse{}, nil
}

// PostApiWaterPumpTransferBase operation middleware

// (POST )
func (s *Server) PostApiWaterPumpTransferBase(ctx context.Context, request PostApiWaterPumpTransferBaseRequestObject) (PostApiWaterPumpTransferBaseResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiWaterPumpTransferBase200JSONResponse{}, nil
}



// GetApiAstanaLrtBuilding1 operation middleware

// (GET )
func (s *Server) GetApiAstanaLrtBuilding1(ctx context.Context, request GetApiAstanaLrtBuilding1RequestObject) (GetApiAstanaLrtBuilding1ResponseObject, error) {
    result := &GetApiAstanaLrtBuilding1200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

// GetApiAstanaLrtStreetlights1 operation middleware

// (GET )
func (s *Server) GetApiAstanaLrtStreetlights1(ctx context.Context, request GetApiAstanaLrtStreetlights1RequestObject) (GetApiAstanaLrtStreetlights1ResponseObject, error) {
    result := &GetApiAstanaLrtStreetlights1200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

// GetApiAstanaLrtTl1 operation middleware

// (GET )
func (s *Server) GetApiAstanaLrtTl1(ctx context.Context, request GetApiAstanaLrtTl1RequestObject) (GetApiAstanaLrtTl1ResponseObject, error) {
    result := &GetApiAstanaLrtTl1200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

// GetApiBusinessCenterBuilding1 operation middleware

// (GET )
func (s *Server) GetApiBusinessCenterBuilding1(ctx context.Context, request GetApiBusinessCenterBuilding1RequestObject) (GetApiBusinessCenterBuilding1ResponseObject, error) {
    result := &GetApiBusinessCenterBuilding1200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

// GetApiGovernmentComplexBuilding1 operation middleware

// (GET )
func (s *Server) GetApiGovernmentComplexBuilding1(ctx context.Context, request GetApiGovernmentComplexBuilding1RequestObject) (GetApiGovernmentComplexBuilding1ResponseObject, error) {
    result := &GetApiGovernmentComplexBuilding1200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

// GetApiGovernmentComplexBuilding2 operation middleware

// (GET )
func (s *Server) GetApiGovernmentComplexBuilding2(ctx context.Context, request GetApiGovernmentComplexBuilding2RequestObject) (GetApiGovernmentComplexBuilding2ResponseObject, error) {
    result := &GetApiGovernmentComplexBuilding2200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

// GetApiGovernmentComplexBuilding3 operation middleware

// (GET )
func (s *Server) GetApiGovernmentComplexBuilding3(ctx context.Context, request GetApiGovernmentComplexBuilding3RequestObject) (GetApiGovernmentComplexBuilding3ResponseObject, error) {
    result := &GetApiGovernmentComplexBuilding3200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

// GetApiGovernmentComplexBuilding4 operation middleware

// (GET )
func (s *Server) GetApiGovernmentComplexBuilding4(ctx context.Context, request GetApiGovernmentComplexBuilding4RequestObject) (GetApiGovernmentComplexBuilding4ResponseObject, error) {
    result := &GetApiGovernmentComplexBuilding4200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

// GetApiHospitalBarrier operation middleware

// (GET )
func (s *Server) GetApiHospitalBarrier(ctx context.Context, request GetApiHospitalBarrierRequestObject) (GetApiHospitalBarrierResponseObject, error) {
    result := &GetApiHospitalBarrier200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

// GetApiHospitalBuilding1 operation middleware

// (GET )
func (s *Server) GetApiHospitalBuilding1(ctx context.Context, request GetApiHospitalBuilding1RequestObject) (GetApiHospitalBuilding1ResponseObject, error) {
    result := &GetApiHospitalBuilding1200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

// GetApiHospitalFire operation middleware

// (GET )
func (s *Server) GetApiHospitalFire(ctx context.Context, request GetApiHospitalFireRequestObject) (GetApiHospitalFireResponseObject, error) {
    result := &GetApiHospitalFire200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

// GetApiHospitalFireAlarm operation middleware

// (GET )
func (s *Server) GetApiHospitalFireAlarm(ctx context.Context, request GetApiHospitalFireAlarmRequestObject) (GetApiHospitalFireAlarmResponseObject, error) {
    result := &GetApiHospitalFireAlarm200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

// GetApiHospitalFireTruck operation middleware

// (GET )
func (s *Server) GetApiHospitalFireTruck(ctx context.Context, request GetApiHospitalFireTruckRequestObject) (GetApiHospitalFireTruckResponseObject, error) {
    result := &GetApiHospitalFireTruck200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

// GetApiHospitalSiren operation middleware

// (GET )
func (s *Server) GetApiHospitalSiren(ctx context.Context, request GetApiHospitalSirenRequestObject) (GetApiHospitalSirenResponseObject, error) {
    result := &GetApiHospitalSiren200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

// GetApiSportsComplexStadiumLights operation middleware

// (GET )
func (s *Server) GetApiSportsComplexStadiumLights(ctx context.Context, request GetApiSportsComplexStadiumLightsRequestObject) (GetApiSportsComplexStadiumLightsResponseObject, error) {
    result := &GetApiSportsComplexStadiumLights200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

// GetApiSportsComplexTl1 operation middleware

// (GET )
func (s *Server) GetApiSportsComplexTl1(ctx context.Context, request GetApiSportsComplexTl1RequestObject) (GetApiSportsComplexTl1ResponseObject, error) {
    result := &GetApiSportsComplexTl1200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

// GetApiWaterPumpBuilding1 operation middleware

// (GET )
func (s *Server) GetApiWaterPumpBuilding1(ctx context.Context, request GetApiWaterPumpBuilding1RequestObject) (GetApiWaterPumpBuilding1ResponseObject, error) {
    result := &GetApiWaterPumpBuilding1200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

// GetApiWaterPumpPump2 operation middleware

// (GET )
func (s *Server) GetApiWaterPumpPump2(ctx context.Context, request GetApiWaterPumpPump2RequestObject) (GetApiWaterPumpPump2ResponseObject, error) {
    result := &GetApiWaterPumpPump2200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

// GetApiWaterPumpPump3 operation middleware

// (GET )
func (s *Server) GetApiWaterPumpPump3(ctx context.Context, request GetApiWaterPumpPump3RequestObject) (GetApiWaterPumpPump3ResponseObject, error) {
    result := &GetApiWaterPumpPump3200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

// GetApiWaterPumpTransferBase operation middleware

// (GET )
func (s *Server) GetApiWaterPumpTransferBase(ctx context.Context, request GetApiWaterPumpTransferBaseRequestObject) (GetApiWaterPumpTransferBaseResponseObject, error) {
    result := &GetApiWaterPumpTransferBase200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

