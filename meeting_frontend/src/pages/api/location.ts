// src/pages/api/location.ts
import { NextApiRequest, NextApiResponse } from 'next';

export default function handler(req: NextApiRequest, res: NextApiResponse) {
  if (req.method === 'POST') {
    const { ip, country, region, city } = req.body;

    // Save the location data to the database or perform other actions
    console.log('Received location data:', { ip, country, region, city });

    res.status(200).json({ message: 'Location data received' });
  } else {
    res.status(405).json({ message: 'Method Not Allowed' });
  }
}
