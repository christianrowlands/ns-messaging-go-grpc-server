package server

import (
	"context"
	pb "craxiom.com/ns-messaging-go-grpc-server/messaging"
	"github.com/rs/zerolog/log"
)

// ConnectionHandshakeServer implements the ConnectionHandshake service
type ConnectionHandshakeServer struct {
	pb.UnimplementedConnectionHandshakeServer
}

func (s *ConnectionHandshakeServer) StartConnection(ctx context.Context, req *pb.ConnectionRequest) (*pb.ConnectionReply, error) {
	log.Info().Msg("Received connection request")
	return &pb.ConnectionReply{ConnectionAccept: true}, nil
}

// WirelessSurveyServer implements the WirelessSurvey service
type WirelessSurveyServer struct {
	pb.UnimplementedWirelessSurveyServer
}

// StreamGsmSurvey handles streaming of GSM survey records
func (s *WirelessSurveyServer) StreamGsmSurvey(stream pb.WirelessSurvey_StreamGsmSurveyServer) error {
	for {
		record, err := stream.Recv()
		if err != nil {
			log.Error().Err(err).Msg("Error receiving GSM record")
			return err
		}
		log.Info().Msgf("Received GSM record: %v", record)
		// Process the record as needed
	}
}

// StreamCdmaSurvey handles streaming of CDMA survey records
func (s *WirelessSurveyServer) StreamCdmaSurvey(stream pb.WirelessSurvey_StreamCdmaSurveyServer) error {
	for {
		record, err := stream.Recv()
		if err != nil {
			log.Error().Err(err).Msg("Error receiving CDMA record")
			return err
		}
		log.Info().Msgf("Received CDMA record: %v", record)
		// Process the record as needed
	}
}

// StreamUmtsSurvey handles streaming of UMTS survey records
func (s *WirelessSurveyServer) StreamUmtsSurvey(stream pb.WirelessSurvey_StreamUmtsSurveyServer) error {
	for {
		record, err := stream.Recv()
		if err != nil {
			log.Error().Err(err).Msg("Error receiving UMTS record")
			return err
		}
		log.Info().Msgf("Received UMTS record: %v", record)
		// Process the record as needed
	}
}

// StreamLteSurvey handles streaming of LTE survey records
func (s *WirelessSurveyServer) StreamLteSurvey(stream pb.WirelessSurvey_StreamLteSurveyServer) error {
	for {
		record, err := stream.Recv()
		if err != nil {
			log.Error().Err(err).Msg("Error receiving LTE record")
			return err
		}
		log.Info().Msgf("Received LTE record: %v", record)
		// Process the record as needed
	}
}

// StreamNrSurvey handles streaming of NR survey records
func (s *WirelessSurveyServer) StreamNrSurvey(stream pb.WirelessSurvey_StreamNrSurveyServer) error {
	for {
		record, err := stream.Recv()
		if err != nil {
			log.Error().Err(err).Msg("Error receiving NR record")
			return err
		}
		log.Info().Msgf("Received NR record: %v", record)
		// Process the record as needed
	}
}

// StreamWifiBeaconSurvey handles streaming of WiFi beacon survey records
func (s *WirelessSurveyServer) StreamWifiBeaconSurvey(stream pb.WirelessSurvey_StreamWifiBeaconSurveyServer) error {
	for {
		record, err := stream.Recv()
		if err != nil {
			log.Error().Err(err).Msg("Error receiving WiFi Beacon record")
			return err
		}
		log.Info().Msgf("Received WiFi Beacon record: %v", record)
		// Process the record as needed
	}
}

// StreamWifiProbeRequestSurvey handles streaming of WiFi probe request survey records
func (s *WirelessSurveyServer) StreamWifiProbeRequestSurvey(stream pb.WirelessSurvey_StreamWifiProbeRequestSurveyServer) error {
	for {
		record, err := stream.Recv()
		if err != nil {
			log.Error().Err(err).Msg("Error receiving WiFi Probe Request record")
			return err
		}
		log.Info().Msgf("Received WiFi Probe Request record: %v", record)
		// Process the record as needed
	}
}

// StreamWifiOtaSurvey handles streaming of WiFi OTA survey records
func (s *WirelessSurveyServer) StreamWifiOtaSurvey(stream pb.WirelessSurvey_StreamWifiOtaSurveyServer) error {
	for {
		record, err := stream.Recv()
		if err != nil {
			log.Error().Err(err).Msg("Error receiving WiFi OTA record")
			return err
		}
		log.Info().Msgf("Received WiFi OTA record: %v", record)
		// Process the record as needed
	}
}

// StreamGnssSurvey handles streaming of GNSS survey records
func (s *WirelessSurveyServer) StreamGnssSurvey(stream pb.WirelessSurvey_StreamGnssSurveyServer) error {
	for {
		record, err := stream.Recv()
		if err != nil {
			log.Error().Err(err).Msg("Error receiving GNSS record")
			return err
		}
		log.Info().Msgf("Received GNSS record: %v", record)
		// Process the record as needed
	}
}

// StreamEnergyDetections handles streaming of energy detection records
func (s *WirelessSurveyServer) StreamEnergyDetections(stream pb.WirelessSurvey_StreamEnergyDetectionsServer) error {
	for {
		record, err := stream.Recv()
		if err != nil {
			log.Error().Err(err).Msg("Error receiving Energy Detection record")
			return err
		}
		log.Info().Msgf("Received Energy Detection record: %v", record)
		// Process the record as needed
	}
}

// StreamSignalDetections handles streaming of signal detection records
func (s *WirelessSurveyServer) StreamSignalDetections(stream pb.WirelessSurvey_StreamSignalDetectionsServer) error {
	for {
		record, err := stream.Recv()
		if err != nil {
			log.Error().Err(err).Msg("Error receiving Signal Detection record")
			return err
		}
		log.Info().Msgf("Received Signal Detection record: %v", record)
		// Process the record as needed
	}
}

// StreamLteRrc handles streaming of LTE RRC records
func (s *WirelessSurveyServer) StreamLteRrc(stream pb.WirelessSurvey_StreamLteRrcServer) error {
	for {
		record, err := stream.Recv()
		if err != nil {
			log.Error().Err(err).Msg("Error receiving LTE RRC record")
			return err
		}
		log.Info().Msgf("Received LTE RRC record: %v", record)
		// Process the record as needed
	}
}

// StreamLteNas handles streaming of LTE NAS records
func (s *WirelessSurveyServer) StreamLteNas(stream pb.WirelessSurvey_StreamLteNasServer) error {
	for {
		record, err := stream.Recv()
		if err != nil {
			log.Error().Err(err).Msg("Error receiving LTE NAS record")
			return err
		}
		log.Info().Msgf("Received LTE NAS record: %v", record)
		// Process the record as needed
	}
}

// DeviceStatusServer implements the DeviceStatus service
type DeviceStatusServer struct {
	pb.UnimplementedDeviceStatusServer
}

func (s *DeviceStatusServer) StatusUpdate(stream pb.DeviceStatus_StatusUpdateServer) error {
	for {
		status, err := stream.Recv()
		if err != nil {
			log.Error().Err(err).Msg("Error receiving device status update")
			return err
		}
		log.Info().Msgf("Received device status: %v", status)
		// Process the status update as needed
	}
	return stream.SendAndClose(&pb.StatusUpdateReply{})
}
