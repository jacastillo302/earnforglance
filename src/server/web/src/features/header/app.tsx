import type { JSX } from "react"
import Localization from "../localizations/app"

export const Header = (): JSX.Element | null => {
  return (

      <header className="header-container">
      {/* Header content goes here */}
      <h1 className='section-header-title-text'>Header</h1>
      <div className="header-localization-container">
        <Localization />
      </div>
    </header>
  )
}