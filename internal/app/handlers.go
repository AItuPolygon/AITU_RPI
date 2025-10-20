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

// PostApiAstanaLrtRgbAlert operation middleware

// (POST )
func (s *Server) PostApiAstanaLrtRgbAlert(ctx context.Context, request PostApiAstanaLrtRgbAlertRequestObject) (PostApiAstanaLrtRgbAlertResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiAstanaLrtRgbAlert200JSONResponse{}, nil
}

// PostApiAstanaLrtRgbB operation middleware

// (POST )
func (s *Server) PostApiAstanaLrtRgbB(ctx context.Context, request PostApiAstanaLrtRgbBRequestObject) (PostApiAstanaLrtRgbBResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiAstanaLrtRgbB200JSONResponse{}, nil
}

// PostApiAstanaLrtRgbG operation middleware

// (POST )
func (s *Server) PostApiAstanaLrtRgbG(ctx context.Context, request PostApiAstanaLrtRgbGRequestObject) (PostApiAstanaLrtRgbGResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiAstanaLrtRgbG200JSONResponse{}, nil
}

// PostApiAstanaLrtRgbR operation middleware

// (POST )
func (s *Server) PostApiAstanaLrtRgbR(ctx context.Context, request PostApiAstanaLrtRgbRRequestObject) (PostApiAstanaLrtRgbRResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiAstanaLrtRgbR200JSONResponse{}, nil
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

// PostApiAstanaLrtTl0 operation middleware

// (POST )
func (s *Server) PostApiAstanaLrtTl0(ctx context.Context, request PostApiAstanaLrtTl0RequestObject) (PostApiAstanaLrtTl0ResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiAstanaLrtTl0200JSONResponse{}, nil
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

// PostApiBusinessCenterRgbAlert operation middleware

// (POST )
func (s *Server) PostApiBusinessCenterRgbAlert(ctx context.Context, request PostApiBusinessCenterRgbAlertRequestObject) (PostApiBusinessCenterRgbAlertResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiBusinessCenterRgbAlert200JSONResponse{}, nil
}

// PostApiBusinessCenterRgbB operation middleware

// (POST )
func (s *Server) PostApiBusinessCenterRgbB(ctx context.Context, request PostApiBusinessCenterRgbBRequestObject) (PostApiBusinessCenterRgbBResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiBusinessCenterRgbB200JSONResponse{}, nil
}

// PostApiBusinessCenterRgbG operation middleware

// (POST )
func (s *Server) PostApiBusinessCenterRgbG(ctx context.Context, request PostApiBusinessCenterRgbGRequestObject) (PostApiBusinessCenterRgbGResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiBusinessCenterRgbG200JSONResponse{}, nil
}

// PostApiBusinessCenterRgbR operation middleware

// (POST )
func (s *Server) PostApiBusinessCenterRgbR(ctx context.Context, request PostApiBusinessCenterRgbRRequestObject) (PostApiBusinessCenterRgbRResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiBusinessCenterRgbR200JSONResponse{}, nil
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

// PostApiGovernmentComplexBuilding5 operation middleware

// (POST )
func (s *Server) PostApiGovernmentComplexBuilding5(ctx context.Context, request PostApiGovernmentComplexBuilding5RequestObject) (PostApiGovernmentComplexBuilding5ResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiGovernmentComplexBuilding5200JSONResponse{}, nil
}

// PostApiGovernmentComplexBuilding6 operation middleware

// (POST )
func (s *Server) PostApiGovernmentComplexBuilding6(ctx context.Context, request PostApiGovernmentComplexBuilding6RequestObject) (PostApiGovernmentComplexBuilding6ResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiGovernmentComplexBuilding6200JSONResponse{}, nil
}

// PostApiGovernmentComplexBuilding7 operation middleware

// (POST )
func (s *Server) PostApiGovernmentComplexBuilding7(ctx context.Context, request PostApiGovernmentComplexBuilding7RequestObject) (PostApiGovernmentComplexBuilding7ResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiGovernmentComplexBuilding7200JSONResponse{}, nil
}

// PostApiGovernmentComplexBuilding8 operation middleware

// (POST )
func (s *Server) PostApiGovernmentComplexBuilding8(ctx context.Context, request PostApiGovernmentComplexBuilding8RequestObject) (PostApiGovernmentComplexBuilding8ResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiGovernmentComplexBuilding8200JSONResponse{}, nil
}

// PostApiGovernmentComplexRgbAlert operation middleware

// (POST )
func (s *Server) PostApiGovernmentComplexRgbAlert(ctx context.Context, request PostApiGovernmentComplexRgbAlertRequestObject) (PostApiGovernmentComplexRgbAlertResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiGovernmentComplexRgbAlert200JSONResponse{}, nil
}

// PostApiGovernmentComplexRgbB operation middleware

// (POST )
func (s *Server) PostApiGovernmentComplexRgbB(ctx context.Context, request PostApiGovernmentComplexRgbBRequestObject) (PostApiGovernmentComplexRgbBResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiGovernmentComplexRgbB200JSONResponse{}, nil
}

// PostApiGovernmentComplexRgbG operation middleware

// (POST )
func (s *Server) PostApiGovernmentComplexRgbG(ctx context.Context, request PostApiGovernmentComplexRgbGRequestObject) (PostApiGovernmentComplexRgbGResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiGovernmentComplexRgbG200JSONResponse{}, nil
}

// PostApiGovernmentComplexRgbR operation middleware

// (POST )
func (s *Server) PostApiGovernmentComplexRgbR(ctx context.Context, request PostApiGovernmentComplexRgbRRequestObject) (PostApiGovernmentComplexRgbRResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiGovernmentComplexRgbR200JSONResponse{}, nil
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

// PostApiHospitalRgbAlert operation middleware

// (POST )
func (s *Server) PostApiHospitalRgbAlert(ctx context.Context, request PostApiHospitalRgbAlertRequestObject) (PostApiHospitalRgbAlertResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiHospitalRgbAlert200JSONResponse{}, nil
}

// PostApiHospitalRgbB operation middleware

// (POST )
func (s *Server) PostApiHospitalRgbB(ctx context.Context, request PostApiHospitalRgbBRequestObject) (PostApiHospitalRgbBResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiHospitalRgbB200JSONResponse{}, nil
}

// PostApiHospitalRgbG operation middleware

// (POST )
func (s *Server) PostApiHospitalRgbG(ctx context.Context, request PostApiHospitalRgbGRequestObject) (PostApiHospitalRgbGResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiHospitalRgbG200JSONResponse{}, nil
}

// PostApiHospitalRgbR operation middleware

// (POST )
func (s *Server) PostApiHospitalRgbR(ctx context.Context, request PostApiHospitalRgbRRequestObject) (PostApiHospitalRgbRResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiHospitalRgbR200JSONResponse{}, nil
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

// PostApiSportsComplexRgbAlert operation middleware

// (POST )
func (s *Server) PostApiSportsComplexRgbAlert(ctx context.Context, request PostApiSportsComplexRgbAlertRequestObject) (PostApiSportsComplexRgbAlertResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiSportsComplexRgbAlert200JSONResponse{}, nil
}

// PostApiSportsComplexRgbB operation middleware

// (POST )
func (s *Server) PostApiSportsComplexRgbB(ctx context.Context, request PostApiSportsComplexRgbBRequestObject) (PostApiSportsComplexRgbBResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiSportsComplexRgbB200JSONResponse{}, nil
}

// PostApiSportsComplexRgbG operation middleware

// (POST )
func (s *Server) PostApiSportsComplexRgbG(ctx context.Context, request PostApiSportsComplexRgbGRequestObject) (PostApiSportsComplexRgbGResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiSportsComplexRgbG200JSONResponse{}, nil
}

// PostApiSportsComplexRgbR operation middleware

// (POST )
func (s *Server) PostApiSportsComplexRgbR(ctx context.Context, request PostApiSportsComplexRgbRRequestObject) (PostApiSportsComplexRgbRResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiSportsComplexRgbR200JSONResponse{}, nil
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

// PostApiSportsComplexStreetlights1 operation middleware

// (POST )
func (s *Server) PostApiSportsComplexStreetlights1(ctx context.Context, request PostApiSportsComplexStreetlights1RequestObject) (PostApiSportsComplexStreetlights1ResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiSportsComplexStreetlights1200JSONResponse{}, nil
}

// PostApiSportsComplexTl0 operation middleware

// (POST )
func (s *Server) PostApiSportsComplexTl0(ctx context.Context, request PostApiSportsComplexTl0RequestObject) (PostApiSportsComplexTl0ResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiSportsComplexTl0200JSONResponse{}, nil
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

// PostApiWaterPumpRgbAlert operation middleware

// (POST )
func (s *Server) PostApiWaterPumpRgbAlert(ctx context.Context, request PostApiWaterPumpRgbAlertRequestObject) (PostApiWaterPumpRgbAlertResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiWaterPumpRgbAlert200JSONResponse{}, nil
}

// PostApiWaterPumpRgbB operation middleware

// (POST )
func (s *Server) PostApiWaterPumpRgbB(ctx context.Context, request PostApiWaterPumpRgbBRequestObject) (PostApiWaterPumpRgbBResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiWaterPumpRgbB200JSONResponse{}, nil
}

// PostApiWaterPumpRgbG operation middleware

// (POST )
func (s *Server) PostApiWaterPumpRgbG(ctx context.Context, request PostApiWaterPumpRgbGRequestObject) (PostApiWaterPumpRgbGResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiWaterPumpRgbG200JSONResponse{}, nil
}

// PostApiWaterPumpRgbR operation middleware

// (POST )
func (s *Server) PostApiWaterPumpRgbR(ctx context.Context, request PostApiWaterPumpRgbRRequestObject) (PostApiWaterPumpRgbRResponseObject, error) {
    s.app.processCommand(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        *request.Body,
    )
    return PostApiWaterPumpRgbR200JSONResponse{}, nil
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

// GetApiAstanaLrtRgbAlert operation middleware

// (GET )
func (s *Server) GetApiAstanaLrtRgbAlert(ctx context.Context, request GetApiAstanaLrtRgbAlertRequestObject) (GetApiAstanaLrtRgbAlertResponseObject, error) {
    result := &GetApiAstanaLrtRgbAlert200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

// GetApiAstanaLrtRgbB operation middleware

// (GET )
func (s *Server) GetApiAstanaLrtRgbB(ctx context.Context, request GetApiAstanaLrtRgbBRequestObject) (GetApiAstanaLrtRgbBResponseObject, error) {
    result := &GetApiAstanaLrtRgbB200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

// GetApiAstanaLrtRgbG operation middleware

// (GET )
func (s *Server) GetApiAstanaLrtRgbG(ctx context.Context, request GetApiAstanaLrtRgbGRequestObject) (GetApiAstanaLrtRgbGResponseObject, error) {
    result := &GetApiAstanaLrtRgbG200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

// GetApiAstanaLrtRgbR operation middleware

// (GET )
func (s *Server) GetApiAstanaLrtRgbR(ctx context.Context, request GetApiAstanaLrtRgbRRequestObject) (GetApiAstanaLrtRgbRResponseObject, error) {
    result := &GetApiAstanaLrtRgbR200JSONResponse{}
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

// GetApiAstanaLrtTl0 operation middleware

// (GET )
func (s *Server) GetApiAstanaLrtTl0(ctx context.Context, request GetApiAstanaLrtTl0RequestObject) (GetApiAstanaLrtTl0ResponseObject, error) {
    result := &GetApiAstanaLrtTl0200JSONResponse{}
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

// GetApiBusinessCenterRgbAlert operation middleware

// (GET )
func (s *Server) GetApiBusinessCenterRgbAlert(ctx context.Context, request GetApiBusinessCenterRgbAlertRequestObject) (GetApiBusinessCenterRgbAlertResponseObject, error) {
    result := &GetApiBusinessCenterRgbAlert200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

// GetApiBusinessCenterRgbB operation middleware

// (GET )
func (s *Server) GetApiBusinessCenterRgbB(ctx context.Context, request GetApiBusinessCenterRgbBRequestObject) (GetApiBusinessCenterRgbBResponseObject, error) {
    result := &GetApiBusinessCenterRgbB200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

// GetApiBusinessCenterRgbG operation middleware

// (GET )
func (s *Server) GetApiBusinessCenterRgbG(ctx context.Context, request GetApiBusinessCenterRgbGRequestObject) (GetApiBusinessCenterRgbGResponseObject, error) {
    result := &GetApiBusinessCenterRgbG200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

// GetApiBusinessCenterRgbR operation middleware

// (GET )
func (s *Server) GetApiBusinessCenterRgbR(ctx context.Context, request GetApiBusinessCenterRgbRRequestObject) (GetApiBusinessCenterRgbRResponseObject, error) {
    result := &GetApiBusinessCenterRgbR200JSONResponse{}
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

// GetApiGovernmentComplexBuilding5 operation middleware

// (GET )
func (s *Server) GetApiGovernmentComplexBuilding5(ctx context.Context, request GetApiGovernmentComplexBuilding5RequestObject) (GetApiGovernmentComplexBuilding5ResponseObject, error) {
    result := &GetApiGovernmentComplexBuilding5200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

// GetApiGovernmentComplexBuilding6 operation middleware

// (GET )
func (s *Server) GetApiGovernmentComplexBuilding6(ctx context.Context, request GetApiGovernmentComplexBuilding6RequestObject) (GetApiGovernmentComplexBuilding6ResponseObject, error) {
    result := &GetApiGovernmentComplexBuilding6200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

// GetApiGovernmentComplexBuilding7 operation middleware

// (GET )
func (s *Server) GetApiGovernmentComplexBuilding7(ctx context.Context, request GetApiGovernmentComplexBuilding7RequestObject) (GetApiGovernmentComplexBuilding7ResponseObject, error) {
    result := &GetApiGovernmentComplexBuilding7200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

// GetApiGovernmentComplexBuilding8 operation middleware

// (GET )
func (s *Server) GetApiGovernmentComplexBuilding8(ctx context.Context, request GetApiGovernmentComplexBuilding8RequestObject) (GetApiGovernmentComplexBuilding8ResponseObject, error) {
    result := &GetApiGovernmentComplexBuilding8200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

// GetApiGovernmentComplexRgbAlert operation middleware

// (GET )
func (s *Server) GetApiGovernmentComplexRgbAlert(ctx context.Context, request GetApiGovernmentComplexRgbAlertRequestObject) (GetApiGovernmentComplexRgbAlertResponseObject, error) {
    result := &GetApiGovernmentComplexRgbAlert200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

// GetApiGovernmentComplexRgbB operation middleware

// (GET )
func (s *Server) GetApiGovernmentComplexRgbB(ctx context.Context, request GetApiGovernmentComplexRgbBRequestObject) (GetApiGovernmentComplexRgbBResponseObject, error) {
    result := &GetApiGovernmentComplexRgbB200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

// GetApiGovernmentComplexRgbG operation middleware

// (GET )
func (s *Server) GetApiGovernmentComplexRgbG(ctx context.Context, request GetApiGovernmentComplexRgbGRequestObject) (GetApiGovernmentComplexRgbGResponseObject, error) {
    result := &GetApiGovernmentComplexRgbG200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

// GetApiGovernmentComplexRgbR operation middleware

// (GET )
func (s *Server) GetApiGovernmentComplexRgbR(ctx context.Context, request GetApiGovernmentComplexRgbRRequestObject) (GetApiGovernmentComplexRgbRResponseObject, error) {
    result := &GetApiGovernmentComplexRgbR200JSONResponse{}
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

// GetApiHospitalRgbAlert operation middleware

// (GET )
func (s *Server) GetApiHospitalRgbAlert(ctx context.Context, request GetApiHospitalRgbAlertRequestObject) (GetApiHospitalRgbAlertResponseObject, error) {
    result := &GetApiHospitalRgbAlert200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

// GetApiHospitalRgbB operation middleware

// (GET )
func (s *Server) GetApiHospitalRgbB(ctx context.Context, request GetApiHospitalRgbBRequestObject) (GetApiHospitalRgbBResponseObject, error) {
    result := &GetApiHospitalRgbB200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

// GetApiHospitalRgbG operation middleware

// (GET )
func (s *Server) GetApiHospitalRgbG(ctx context.Context, request GetApiHospitalRgbGRequestObject) (GetApiHospitalRgbGResponseObject, error) {
    result := &GetApiHospitalRgbG200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

// GetApiHospitalRgbR operation middleware

// (GET )
func (s *Server) GetApiHospitalRgbR(ctx context.Context, request GetApiHospitalRgbRRequestObject) (GetApiHospitalRgbRResponseObject, error) {
    result := &GetApiHospitalRgbR200JSONResponse{}
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

// GetApiSportsComplexRgbAlert operation middleware

// (GET )
func (s *Server) GetApiSportsComplexRgbAlert(ctx context.Context, request GetApiSportsComplexRgbAlertRequestObject) (GetApiSportsComplexRgbAlertResponseObject, error) {
    result := &GetApiSportsComplexRgbAlert200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

// GetApiSportsComplexRgbB operation middleware

// (GET )
func (s *Server) GetApiSportsComplexRgbB(ctx context.Context, request GetApiSportsComplexRgbBRequestObject) (GetApiSportsComplexRgbBResponseObject, error) {
    result := &GetApiSportsComplexRgbB200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

// GetApiSportsComplexRgbG operation middleware

// (GET )
func (s *Server) GetApiSportsComplexRgbG(ctx context.Context, request GetApiSportsComplexRgbGRequestObject) (GetApiSportsComplexRgbGResponseObject, error) {
    result := &GetApiSportsComplexRgbG200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

// GetApiSportsComplexRgbR operation middleware

// (GET )
func (s *Server) GetApiSportsComplexRgbR(ctx context.Context, request GetApiSportsComplexRgbRRequestObject) (GetApiSportsComplexRgbRResponseObject, error) {
    result := &GetApiSportsComplexRgbR200JSONResponse{}
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

// GetApiSportsComplexStreetlights1 operation middleware

// (GET )
func (s *Server) GetApiSportsComplexStreetlights1(ctx context.Context, request GetApiSportsComplexStreetlights1RequestObject) (GetApiSportsComplexStreetlights1ResponseObject, error) {
    result := &GetApiSportsComplexStreetlights1200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

// GetApiSportsComplexTl0 operation middleware

// (GET )
func (s *Server) GetApiSportsComplexTl0(ctx context.Context, request GetApiSportsComplexTl0RequestObject) (GetApiSportsComplexTl0ResponseObject, error) {
    result := &GetApiSportsComplexTl0200JSONResponse{}
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

// GetApiWaterPumpRgbAlert operation middleware

// (GET )
func (s *Server) GetApiWaterPumpRgbAlert(ctx context.Context, request GetApiWaterPumpRgbAlertRequestObject) (GetApiWaterPumpRgbAlertResponseObject, error) {
    result := &GetApiWaterPumpRgbAlert200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

// GetApiWaterPumpRgbB operation middleware

// (GET )
func (s *Server) GetApiWaterPumpRgbB(ctx context.Context, request GetApiWaterPumpRgbBRequestObject) (GetApiWaterPumpRgbBResponseObject, error) {
    result := &GetApiWaterPumpRgbB200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

// GetApiWaterPumpRgbG operation middleware

// (GET )
func (s *Server) GetApiWaterPumpRgbG(ctx context.Context, request GetApiWaterPumpRgbGRequestObject) (GetApiWaterPumpRgbGResponseObject, error) {
    result := &GetApiWaterPumpRgbG200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

// GetApiWaterPumpRgbR operation middleware

// (GET )
func (s *Server) GetApiWaterPumpRgbR(ctx context.Context, request GetApiWaterPumpRgbRRequestObject) (GetApiWaterPumpRgbRResponseObject, error) {
    result := &GetApiWaterPumpRgbR200JSONResponse{}
    s.app.processInfo(
        ctx.Value(TaskKey).(string),
        ctx.Value(ElementKey).(string),
        result,
    )
    return result, nil
}

