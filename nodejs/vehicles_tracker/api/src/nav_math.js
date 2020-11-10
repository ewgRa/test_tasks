function toRadians(deg) { return deg * (Math.PI / 180); }
function toDegrees(rad) { return rad * (180 / Math.PI); }

const R = 637100000; // Radius of the earth in m

exports.distance = function distance(lat1, lng1, lat2, lng2) {
  const dLat = toRadians(lat2 - lat1);
  const dLng = toRadians(lng2 - lng1);

  const a = Math.sin(dLat / 2) * Math.sin(dLat / 2)
        + Math.cos(toRadians(lat1)) * Math.cos(toRadians(lat2))
        * Math.sin(dLng / 2) * Math.sin(dLng / 2);

  return R * 2 * Math.atan2(Math.sqrt(a), Math.sqrt(1 - a));
};

exports.bearing = function bearing(lat1, lng1, lat2, lng2) {
  const phi1 = toRadians(lat1);
  const phi2 = toRadians(lat2);
  const dLng = toRadians(lng2 - lng1);

  const y = Math.sin(dLng) * Math.cos(phi2);
  const x = Math.cos(phi1) * Math.sin(phi2) - Math.sin(phi1) * Math.cos(phi2) * Math.cos(dLng);

  return toDegrees(Math.atan2(y, x));
};
