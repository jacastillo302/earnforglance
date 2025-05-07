import './css/App.css'
import { Counter } from "./features/counter/Counter"
import { Quotes } from "./features/quotes/Quotes"
import { Layaut } from "./features/layout/layout"

export const App = () => (
  <Layaut>
    <div className="app-container">
      <h1 className="section-header-title-text">EarnForGlance</h1>
      <div className="app-content">
        <Counter />
        <Quotes />
      </div>
    </div>
  </Layaut>
)
