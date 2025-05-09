import React from 'react';
import Menu from '@mui/joy/Menu';
import MenuItem from '@mui/joy/MenuItem';
import ListItemDecorator from '@mui/joy/ListItemDecorator';
import AccountCircleIcon from '@mui/icons-material/AccountCircle';
import ListAltIcon from '@mui/icons-material/ListAlt';
import MonetizationOnIcon from '@mui/icons-material/MonetizationOn';
import VisibilityIcon from '@mui/icons-material/Visibility';
import PlaylistAddCheckIcon from '@mui/icons-material/PlaylistAddCheck';
import LogoutIcon from '@mui/icons-material/Logout';
import { Link } from '@mui/joy';

interface UserMenuProps {
  anchorEl: HTMLElement | null;
  open: boolean;
  onClose: () => void;
  onLogout: () => void;
}

const UserMenu: React.FC<UserMenuProps> = ({ anchorEl, open, onClose, onLogout }) => {
  const menuItems = [
    { label: "Update Account", href: "/account", icon: <AccountCircleIcon /> },
    { label: "Orders", href: "/orders", icon: <ListAltIcon /> },
    { label: "Earn For Glance Affiliate", href: "/affiliate", icon: <MonetizationOnIcon /> },
    { label: "Watch List", href: "/watchlist", icon: <VisibilityIcon /> },
    { label: "My List", href: "/mylist", icon: <PlaylistAddCheckIcon /> },
  ];

  return (
    <Menu
      anchorEl={anchorEl}
      open={open}
      onClose={onClose}
      placement="bottom-end"
      sx={{ minWidth: 200, zIndex: 1301 }} // Ensure menu is above other elements
    >
      {menuItems.map((item) => (
        <MenuItem key={item.label} onClick={onClose} component={Link} href={item.href} sx={{textDecoration: 'none', color: 'inherit'}}>
          <ListItemDecorator>{item.icon}</ListItemDecorator>
          {item.label}
        </MenuItem>
      ))}
      <MenuItem onClick={() => { onLogout(); onClose(); }}>
        <ListItemDecorator><LogoutIcon /></ListItemDecorator>
        Logout
      </MenuItem>
    </Menu>
  );
};

export default UserMenu;
