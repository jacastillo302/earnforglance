import React from 'react';
import HeaderContainer from './HeaderContainer';
import SideNavContainer from './SideNavContainer';
import BodyContainer from './BodyContainer';
import FooterContainer from './FooterContainer';

interface LayoutProps {
  children?: React.ReactNode;
}

const Layout: React.FC<LayoutProps> = ({ children }) => {
  return (
    <div className="mms-body">
      <HeaderContainer />
      <div className="main-content">
        <SideNavContainer />
        <BodyContainer>{children}</BodyContainer>
      </div>
      <FooterContainer />
    </div>
  );
};

export default Layout;