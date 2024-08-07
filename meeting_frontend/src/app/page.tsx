"use client";

import React from 'react';
import GeoLocation from '@/components/GeoLocation';
import VideoChat from '@/components/VideoChat';

const HomePage = () => {
  return (
    <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
      <div className="bg-white p-4 shadow rounded">
        <h2 className="text-xl font-bold mb-4">Your Location</h2>
        <GeoLocation />
      </div>
      <div className="bg-white p-4 shadow rounded">
        <h2 className="text-xl font-bold mb-4">Video Chat</h2>
        <VideoChat />
      </div>
    </div>
  );
};

export default HomePage;
