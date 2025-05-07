import type { JSX } from "react"
import { Footer } from './../footer/footer';
import { Header } from './../header/header';
import { Navbar } from './../navbar/navbar';
import { Content } from './../content/content';

interface LayoutProps {
  children?: React.ReactNode;
}

export const Layaut = ({ children }: LayoutProps): JSX.Element | null => {
  return (
    <div className="mms-body">
      <Header />
      <div className="main-content">
        <Navbar />
        <Content>{children}</Content>
      </div>
      <Footer />
    </div>
  )
}