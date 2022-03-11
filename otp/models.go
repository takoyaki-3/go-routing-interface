package otp

type OTP struct {
	Plan struct {
		Date int64 `json:"date"`
		From struct {
			Name       string  `json:"name"`
			Lon        float64 `json:"lon"`
			Lat        float64 `json:"lat"`
			VertexType string  `json:"vertexType"`
		} `json:"from"`
		To struct {
			Name       string  `json:"name"`
			Lon        float64 `json:"lon"`
			Lat        float64 `json:"lat"`
			VertexType string  `json:"vertexType"`
		} `json:"to"`
		Itineraries []struct {
			Duration          int     `json:"duration"`
			StartTime         int64   `json:"startTime"`
			EndTime           int64   `json:"endTime"`
			WalkTime          int     `json:"walkTime"`
			TransitTime       int     `json:"transitTime"`
			WaitingTime       int     `json:"waitingTime"`
			WalkDistance      float64 `json:"walkDistance"`
			WalkLimitExceeded bool    `json:"walkLimitExceeded"`
			ElevationLost     float64 `json:"elevationLost"`
			ElevationGained   float64 `json:"elevationGained"`
			Transfers         int     `json:"transfers"`
			Legs              []struct {
				StartTime                int64   `json:"startTime"`
				EndTime                  int64   `json:"endTime"`
				DepartureDelay           int     `json:"departureDelay"`
				ArrivalDelay             int     `json:"arrivalDelay"`
				RealTime                 bool    `json:"realTime"`
				Distance                 float64 `json:"distance"`
				Pathway                  bool    `json:"pathway"`
				Mode                     string  `json:"mode"`
				TransitLeg               bool    `json:"transitLeg"`
				Route                    string  `json:"route,omitempty"`
				AgencyTimeZoneOffset     int     `json:"agencyTimeZoneOffset"`
				InterlineWithPreviousLeg bool    `json:"interlineWithPreviousLeg,omitempty"`
				To                       struct {
					Name       string  `json:"name"`
					StopID     string  `json:"stopId"`
					Lon        float64 `json:"lon"`
					Lat        float64 `json:"lat"`
					Arrival    int64   `json:"arrival"`
					Departure  int64   `json:"departure"`
					ZoneID     string  `json:"zoneId"`
					VertexType string  `json:"vertexType"`
				} `json:"to,omitempty"`
				LegGeometry struct {
					Points string `json:"points"`
					Length int    `json:"length"`
				} `json:"legGeometry"`
				Steps []struct {
					Distance          float64 `json:"distance"`
					RelativeDirection string  `json:"relativeDirection"`
					StreetName        string  `json:"streetName"`
					AbsoluteDirection string  `json:"absoluteDirection"`
					StayOn            bool    `json:"stayOn"`
					Area              bool    `json:"area"`
					BogusName         bool    `json:"bogusName"`
					Lon               float64 `json:"lon"`
					Lat               float64 `json:"lat"`
					Elevation         string  `json:"elevation"`
				} `json:"steps"`
				RentedBike     bool    `json:"rentedBike,omitempty"`
				Duration       float64 `json:"duration"`
				AgencyName     string  `json:"agencyName,omitempty"`
				AgencyURL      string  `json:"agencyUrl,omitempty"`
				RouteID        string  `json:"routeId,omitempty"`
				Headsign       string  `json:"headsign,omitempty"`
				AgencyID       string  `json:"agencyId,omitempty"`
				TripID         string  `json:"tripId,omitempty"`
				ServiceDate    string  `json:"serviceDate,omitempty"`
				RouteShortName string  `json:"routeShortName,omitempty"`
				From           struct {
					Name       string  `json:"name"`
					StopID     string  `json:"stopId"`
					Lon        float64 `json:"lon"`
					Lat        float64 `json:"lat"`
					Arrival    int64   `json:"arrival"`
					Departure  int64   `json:"departure"`
					ZoneID     string  `json:"zoneId"`
					VertexType string  `json:"vertexType"`
				} `json:"from,omitempty"`
			} `json:"legs"`
			TooSloped bool `json:"tooSloped"`
		} `json:"itineraries"`
	} `json:"plan"`
}
