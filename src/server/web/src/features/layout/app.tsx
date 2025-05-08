import type { JSX } from "react"
import { Footer } from '../footer/footer';
import { Header } from '../header/app';
import { Navbar } from '../navbar/sidebar';
import { Content } from '../content/content';
import Box from '@mui/joy/Box';

interface LayoutProps {
  children?: React.ReactNode;
}

export const Layaut = ({ children }: LayoutProps): JSX.Element | null => {
  return (
    <Box sx={{ display: 'flex', flexDirection: 'column', minHeight: '100vh' }}>
      {/* Header at the top */}
      <Box component="header">
        <Header />
      </Box>

      {/* Main content area: Navbar on left, Content on right/middle */}
      <Box
        sx={{
          display: 'flex',
          flexGrow: 1,
          flexDirection: { xs: 'column', md: 'row' }, // Stack on small screens, row on medium+
        }}
      >
        {/* Content area taking most space */}
        <Box
          component="main"
          sx={{
            flexGrow: 1,
            p: { xs: 1, sm: 2 }, // Padding responsive to screen size
            order: { xs: 2, md: 2 }, // Content below navbar on xs, right of navbar on md
            borderRight: { xs: 'none', md: '1px solid' }, // Border on the right of content
            borderColor: 'divider', // Border color
          }}
        >
          <Content>{children}</Content>
        </Box>

        {/* Navbar on the left */}
        <Box
          component="aside" // Using aside for sidebar/navbar
          sx={{
            width: { xs: '100%', md: 250 }, // Full width on xs, fixed width on md
            flexShrink: 0, // Prevent navbar from shrinking on flex layouts
            p: { xs: 1, sm: 2 }, // Padding responsive to screen size
            order: { xs: 1, md: 1 }, // Navbar above content on xs, left of content on md
            // Example: backgroundColor: 'background.level1', // Optional: for visual distinction
          }}
        >
          <Navbar />
        </Box>
      </Box>

      {/* Footer at the bottom */}
      <Box component="footer">
        <Footer />
      </Box>
    </Box>
  )
}