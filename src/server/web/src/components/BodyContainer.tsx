import React from 'react';

interface BodyContainerProps {
  children?: React.ReactNode; // Allow children to be passed
}

const BodyContainer: React.FC<BodyContainerProps> = ({ children }) => {
  return (
    <main className="body-container">
      {/* Main application content goes here */}
      <h1>EarnForGlance</h1>
      {children}
    </main>
  );
};

export default BodyContainer;
