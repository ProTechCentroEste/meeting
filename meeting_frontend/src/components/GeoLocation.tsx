// src/components/GeoLocation.tsx
// APIs like ipinfo.io, ipstack, or ipgeolocation.io.
// src/components/GeoLocation.tsx (updated)
import React, { useEffect, useState } from 'react';

const GeoLocation = () => {
  const [location, setLocation] = useState(null);

  useEffect(() => {
    const fetchLocation = async () => {
      try {
        const response = await fetch('https://ipinfo.io?token=YOUR_API_TOKEN');
        const data = await response.json();
        setLocation(data);
        
        // Send location data to the backend
        await fetch('/api/location', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify(data)
        });
      } catch (error) {
        console.error('Error fetching geolocation:', error);
      }
    };

    fetchLocation();
  }, []);

  return (
    <div>
      {location ? (
        <div>
          <p>IP: {location.ip}</p>
          <p>Country: {location.country}</p>
          <p>Region: {location.region}</p>
          <p>City: {location.city}</p>
        </div>
      ) : (
        <p>Loading...</p>
      )}
    </div>
  );
};

export default GeoLocation;
