import React from 'react'
import { Map, Marker, Popup, TileLayer } from 'react-leaflet'
import { w3cwebsocket as W3CWebSocket } from "websocket";

const position = [52.53, 13.403];

class App extends React.Component {
    constructor() {
        super();
        this.state = {
            markers: []
        };
    }

    deleteMarker = (id) => {
        let {markers} = this.state;
        markers = markers.filter(marker => marker.id !== id);
        this.setState({markers});
    };

    updateMarker = (id, lat, lng) => {
        const {markers} = this.state;

        const exists = markers.find(function (elem) {
            return elem.id === id;
        });

        if (exists) {
            exists.position = [lat, lng];
        } else {
            markers.push({id: id, position: [lat, lng]});
        }

        this.setState({markers})
    };

    componentWillMount() {
    const client = new W3CWebSocket(process.env.REACT_APP_BROADCAST_URL);

    client.onopen = () => {
      console.log('WebSocket Client Connected');
    };
    client.onmessage = (message) => {
        const data = JSON.parse(message.data);

        switch (data.type) {
            case "location_update":
                this.updateMarker(data.id, data.lat, data.lng);
                break;
            case "de-register":
                this.deleteMarker(data.id);
                break;
            default:
        }
    };
  }

  render() {
    return (
        // Important! Always set the container height explicitly
        <div style={{ height: '400px', width: '100%' }}>
            <Map center={position} zoom={13} onClick={this.addMarker}>
            <TileLayer
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
        attribution="&copy; <a href=&quot;http://osm.org/copyright&quot;>OpenStreetMap</a> contributors"
            />
                {this.state.markers.map((data, idx) =>
                    <Marker key={`marker-${idx}`} position={data.position}>
                        <Popup>
                            <span>Car {data.id}</span>
                        </Popup>
                    </Marker>
                )}
        </Map>
        </div>
    );
  }
}

export default App;
