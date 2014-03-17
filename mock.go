package gosubsonic

// mockData maps a mock URL to mock data from the mockTable
var mockData map[string][]byte

// mockTable maps a method to mock JSON data for testing
var mockTable = []struct {
	method string
	data []byte
}{
	{"ping", []byte(`{"subsonic-response":{
		"status": "ok",
		"xmlns": "http://subsonic.org/restapi",
		"version": "1.9.0"
	}}`)},
	{"getLicense", []byte(`{"subsonic-response": {
		"status": "ok",
		"xmlns": "http://subsonic.org/restapi",
		"license": {
			"valid": true,
			"email": "mock@example.com",
			"date": "2014-01-01T00:00:00",
			"key": "abcdef0123456789abcdef0123456789"
		},
		"version": "1.9.0"
	}}`)},
	{"getMusicFolders", []byte(`{"subsonic-response": {
		"status": "ok",
		"xmlns": "http://subsonic.org/restapi",
		"musicFolders": {"musicFolder": {
			"id": 0,
			"name": "Music"
		}},
		"version": "1.9.0"
	}}`)},
	{"getIndexes", []byte(`{"subsonic-response": {
		"status": "ok",
		"indexes": {
			"index": [{
				"name": "A",
				"artist": {
					"id": 1,
					"name": "Adventure"
				}
			},
			{
				"name": "B",
				"artist": {
					"id": 2,
					"name": "Boston"
				}
			}],
			"lastModified": 1395014311154
		},
		"xmlns": "http://subsonic.org/restapi",
		"version": "1.9.0"
	}}`)},
}

// mockInit generates the mock data map, so we can test gosubsonic against known, static data
func mockInit(s Client) error {
	// Initialize map
	mockData = map[string][]byte{}

	// Populate map using this client's URLs
	for _, entry := range mockTable {
		mockData[s.makeURL(entry.method)] = entry.data
	}

	return nil
}
