import type { JSX } from "react"

export const Navbar = (): JSX.Element | null => {
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
  )
}