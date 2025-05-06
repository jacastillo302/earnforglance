import React from 'react';

interface SideNavContainerProps {
  // Define props here if needed
}

const SideNavContainer: React.FC<SideNavContainerProps> = () => {
  return (
    <aside className="side-nav-container">
      {/* Side navigation content goes here */}
      <nav>
        <ul>
          <li>Link 1</li>
          <li>Link 2</li>
        </ul>
      </nav>
    </aside>
  );
};

export default SideNavContainer;
