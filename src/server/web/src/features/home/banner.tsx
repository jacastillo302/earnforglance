import React, { useState, useEffect } from 'react';
import Box from '@mui/joy/Box';
import Sheet from '@mui/joy/Sheet';

const BannerPage: React.FC = () => {
  // Placeholder for images - in a real slider, this would be dynamic
  const images = [
    { src: import.meta.env.VITE_MEDIA_URL + 'images/banner/1.jpg', alt: 'Slide 1' },
    { src: import.meta.env.VITE_MEDIA_URL + 'images/banner/2.jpg', alt: 'Slide 2' },
    { src: import.meta.env.VITE_MEDIA_URL + 'images/banner/3.jpg', alt: 'Slide 3' },
    { src: import.meta.env.VITE_MEDIA_URL + 'images/banner/4.jpg', alt: 'Slide 4' },
  ];

  const [currentImageIndex, setCurrentImageIndex] = useState(0);

  useEffect(() => {
    const timer = setTimeout(() => {
      setCurrentImageIndex((prevIndex) =>
        prevIndex === images.length - 1 ? 0 : prevIndex + 1
      );
    }, 3000); // Change image every 3 seconds

    return () => clearTimeout(timer); // Cleanup the timer on component unmount
  }, [currentImageIndex, images.length]);

  const currentImage = images[currentImageIndex];

  return (
    <Sheet
      variant="outlined" // or "plain", "soft"
      sx={{
        height: '20vh',
        overflow: 'hidden', // To contain the image if it's larger
        position: 'relative', // For potential absolute positioning of controls
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
        bgcolor: 'background.level1', // Example background color
      }}
    >
      {/* This Box would be the content of a single slide */}
      <Box
        component="img"
        src={currentImage.src}
        alt={currentImage.alt}
        sx={{
          width: '100%',
          height: '100%',
          objectFit: 'cover', // Ensures the image covers the area, might crop
        }}
      />
      {/* Slider controls (dots, arrows) would go here */}
    </Sheet>
  );
};

export default BannerPage;
