import React, { useEffect, useRef, useState } from 'react';
import Peer from 'simple-peer';

const VideoChat = () => {
  const [peer, setPeer] = useState<Peer.Instance | null>(null);
  const localVideoRef = useRef<HTMLVideoElement>(null);
  const remoteVideoRef = useRef<HTMLVideoElement>(null);

  useEffect(() => {
    const peer = new Peer({ initiator: true, trickle: false });
    setPeer(peer);

    navigator.mediaDevices.getUserMedia({ video: true, audio: true }).then((stream) => {
      if (localVideoRef.current) {
        localVideoRef.current.srcObject = stream;
      }
      peer.addStream(stream);
    });

    peer.on('stream', (stream) => {
      if (remoteVideoRef.current) {
        remoteVideoRef.current.srcObject = stream;
      }
    });

    return () => {
      peer.destroy();
    };
  }, []);

  return (
    <div>
      <video ref={localVideoRef} autoPlay playsInline muted />
      <video ref={remoteVideoRef} autoPlay playsInline />
    </div>
  );
};

export default VideoChat;
