import type { JSX } from "react";
import { useState } from 'react'; // Import useState
import List from '@mui/joy/List';
import ListItem from '@mui/joy/ListItem';
import ListItemButton from '@mui/joy/ListItemButton';
import ListItemDecorator from '@mui/joy/ListItemDecorator';
import Typography from '@mui/joy/Typography';
import { Home, Apps, Settings, KeyboardArrowDown } from '@mui/icons-material'; // Example icons
import './sidebar.css'; // Import the CSS file
import ListItemContent from "@mui/joy/ListItemContent";

export const Navbar = (): JSX.Element | null => {
  const [openApps, setOpenApps] = useState(false);
  const [openSettings, setOpenSettings] = useState(false);

  return (
    <aside className="side-nav-container" style={{ float: 'left' }}>
      <List className="modern-menu"
        sx={{
          '--List-nestedInsetStart': '40px', // Indent nested items
          // Add this to ensure proper layout with nested items
          '& .MuiListItem-root': {
            flexDirection: 'column',
            alignItems: 'stretch',
          },
          '& .MuiListItemButton-root': {
            display: 'flex',
            justifyContent: 'flex-start',
            width: '100%',
          }
        }}
      >
        <ListItem>
          <ListItemButton>
            <ListItemDecorator>
              <Home />
            </ListItemDecorator>
            <ListItemContent><Typography>Dashboard</Typography></ListItemContent>
          </ListItemButton>
        </ListItem>
        <ListItem nested className={openApps ? 'Mui-expanded' : ''}>
          <ListItemButton onClick={() => setOpenApps(!openApps)}>
            <ListItemDecorator>
              <Apps />
            </ListItemDecorator>
            <ListItemContent><Typography>Applications</Typography></ListItemContent>
            <KeyboardArrowDown sx={{ ml: 'auto', transform: openApps ? 'rotate(180deg)' : 'rotate(0deg)', transition: 'transform 0.3s' }} />
          </ListItemButton>
          <List sx={{ '--List-gap': '0px', maxHeight: openApps ? '500px' : '0', opacity: openApps ? 1 : 0, visibility: openApps ? 'visible' : 'hidden', overflow: 'hidden', transition: 'max-height 0.3s ease-in-out, opacity 0.3s ease-in-out, visibility 0.3s ease-in-out' }}>
            <ListItem>
              <ListItemButton><ListItemContent><Typography>App 1</Typography></ListItemContent></ListItemButton>
            </ListItem>
            <ListItem>
              <ListItemButton><ListItemContent><Typography>App 2</Typography></ListItemContent></ListItemButton>
            </ListItem>
          </List>
        </ListItem>
        <ListItem nested className={openSettings ? 'Mui-expanded' : ''}>
          <ListItemButton onClick={() => setOpenSettings(!openSettings)}>
            <ListItemDecorator>
              <Settings />
            </ListItemDecorator>
            <ListItemContent><Typography>Settings</Typography></ListItemContent>
            <KeyboardArrowDown sx={{ ml: 'auto', transform: openSettings ? 'rotate(180deg)' : 'rotate(0deg)', transition: 'transform 0.3s' }} />
          </ListItemButton>
          <List sx={{ '--List-gap': '0px', maxHeight: openSettings ? '500px' : '0', opacity: openSettings ? 1 : 0, visibility: openSettings ? 'visible' : 'hidden', overflow: 'hidden', transition: 'max-height 0.3s ease-in-out, opacity 0.3s ease-in-out, visibility 0.3s ease-in-out' }}>
            <ListItem>
              <ListItemButton><ListItemContent><Typography>Profile</Typography></ListItemContent></ListItemButton>
            </ListItem>
            <ListItem>
              <ListItemButton><ListItemContent><Typography>Security</Typography></ListItemContent></ListItemButton>
            </ListItem>
          </List>
        </ListItem>
      </List>
    </aside>
  );
};