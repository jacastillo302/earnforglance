import type { JSX } from "react";
import React, { useState, useEffect } from "react";
import ShoppingCartCheckoutIcon from '@mui/icons-material/ShoppingCartCheckout';
import Localization from "../localization/app";
import { Box, Button, IconButton, Link, Sheet, Typography, Drawer, List, ListItem, ListItemButton, Input } from "@mui/joy";
import MenuIcon from "@mui/icons-material/Menu";
import SearchIcon from "@mui/icons-material/Search";
import Avatar from '@mui/material/Avatar';
import Stack from '@mui/material/Stack';
import UserMenu from './usermenu';

export const Header = (): JSX.Element | null => {
  const [open, setOpen] = useState(false);
  const [allMenuOpen, setAllMenuOpen] = useState(false); // New state for "All" menu
  const [isAuthenticated, setIsAuthenticated] = useState(false); // New state for authentication
  const [userMenuAnchorEl, setUserMenuAnchorEl] = useState<null | HTMLElement>(null); // State for UserMenu anchor

  useEffect(() => {
    const checkAuth = () => {
      const token = localStorage.getItem('authToken');
      setIsAuthenticated(!!token);
    };

    checkAuth(); // Initial check

    const handleStorageChange = (event: StorageEvent) => {
      if (event.key === 'authToken') {
        checkAuth();
      }
    };

    window.addEventListener('storage', handleStorageChange);

    return () => {
      window.removeEventListener('storage', handleStorageChange);
    };
  }, []);

  const handleUserMenuOpen = (event: React.MouseEvent<HTMLElement>) => {
    setUserMenuAnchorEl(event.currentTarget);
  };

  const handleUserMenuClose = () => {
    setUserMenuAnchorEl(null);
  };

  const handleLogout = () => {
    localStorage.removeItem('authToken');
    setIsAuthenticated(false);
    handleUserMenuClose(); // Close the menu after logout
  };

  const toggleDrawer = (inOpen: boolean) => (event: React.KeyboardEvent | React.MouseEvent) => {
    if (
      event.type === 'keydown' &&
      ((event as React.KeyboardEvent).key === 'Tab' || (event as React.KeyboardEvent).key === 'Shift')
    ) {
      return;
    }
    setOpen(inOpen);
  };

  const toggleAllMenuDrawer = (inOpen: boolean) => (event: React.KeyboardEvent | React.MouseEvent) => {
    if (
      event.type === 'keydown' &&
      ((event as React.KeyboardEvent).key === 'Tab' || (event as React.KeyboardEvent).key === 'Shift')
    ) {
      return;
    }
    setAllMenuOpen(inOpen);
  };

  const menuItems = [
    { label: "Earn for Glance", href: "/deals" },
    { label: "News", href: "/news" },
    { label: "Blogs", href: "/blogs" },
    { label: "Contact Us", href: "/contact" },
    { label: "Customer Service", href: "/customer-service" },
  ];

  const allMenuItems = [
    { label: "Electronics", href: "/electronics" },
    { label: "Books", href: "/books" },
    { label: "Home & Kitchen", href: "/home-kitchen" },
    { label: "Fashion", href: "/fashion" },
    { label: "Health & Beauty", href: "/health-beauty" },
    // Add more categories as needed
  ];

  return (
    <Sheet component="header">
      <Box sx={{ display: 'inline', gap: 1, alignItems: 'center', width: '100%' }}>
        <Box sx={{ flexGrow: 1, mx: 2, display: { xs: 'none', md: 'flex' }, flexDirection: 'row', backgroundColor: '#131921', marginLeft:0, marginRight:0 }}>
          <Box sx={{ display: 'flex', alignItems: 'center', gap: 1, marginLeft: 2, marginTop: "10px", marginBottom: 1 }}>
            <Typography level="h4" component="h1">
              <a href="/" style={{ textDecoration: 'none', color: 'white' }}>
                <img src="/logo.png" alt="PERCONEX" style={{ height: '55px' }} />
              </a>
            </Typography>
          </Box>
          <Box sx={{ display: 'inline', gap: 1, alignItems: 'center', width: '100%', marginTop: 1.5, marginLeft: '120px' }}>
            <Input
              startDecorator={<SearchIcon />}
              placeholder="Search your products"
              sx={{
                width: '100%',
                maxWidth: '80%', // Constrain search bar width and center it
                height: '2.8rem', // Height of the search bar
                '--Input-focusedThickness': '0rem', // Remove focus ring for a cleaner look
                mb: 0.5, // Margin bottom for spacing before menu items
              }}
            />
          </Box>
          <Box sx={{ display: 'flex', gap: 1, alignItems: 'center' }}>
            <Localization />
            <IconButton
              variant="outlined"
              color="neutral"
              onClick={toggleDrawer(true)}
              sx={{ display: { md: 'none' }, color: 'white', borderColor: '#555' }}
            >
              <MenuIcon />
            </IconButton>
            <Drawer open={allMenuOpen} onClose={toggleAllMenuDrawer(false)}>
            <Box
              role="presentation"
              onClick={toggleAllMenuDrawer(false)}
              onKeyDown={toggleAllMenuDrawer(false)}
              sx={{ minWidth: 250 }}
            >
              <List>
                <ListItem>
                  <Typography level="h4" sx={{ fontWeight: 'bold', p: 1.5 }}>Shop by Department</Typography>
                </ListItem>
                {allMenuItems.map((item) => (
                  <ListItem key={item.label}>
                    <ListItemButton component="a" href={item.href}>
                      {item.label}
                    </ListItemButton>
                  </ListItem>
                ))}
              </List>
            </Box>
            </Drawer>
          {/* Mobile Search Bar - Placed inside the main drawer for mobile view */}
            <Drawer open={open} onClose={toggleDrawer(false)}>
            <Box
              role="presentation"
              sx={{ minWidth: 250, p: 2 }}
              >
              <Box sx={{ mb: 2 }}>
                <Input
                  startDecorator={<SearchIcon />}
                  placeholder="Search YourLogo"
                  sx={{ width: '100%' }}
                  onClick={(e) => e.stopPropagation()} // Prevent drawer from closing on click
                  onKeyDown={(e) => e.stopPropagation()} // Prevent drawer from closing on keydown
                />
              </Box>
              <List
                onClick={toggleDrawer(false)}
                onKeyDown={toggleDrawer(false)}
              >
                {menuItems.map((item) => (
                  <ListItem key={item.label}>
                    <ListItemButton component="a" href={item.href}>
                      {item.label}
                    </ListItemButton>
                  </ListItem>
                ))}
                {!isAuthenticated && (
                  <>
                    <ListItem>
                      <ListItemButton component="a" href="/login">Login</ListItemButton>
                    </ListItem>
                    <ListItem>
                      <ListItemButton component="a" href="/register">Register</ListItemButton>
                    </ListItem>
                  </>
                )}
              </List>
            </Box>
            </Drawer>
          <Box sx={{ display: 'flex', gap: 1, alignItems: 'center', marginRight: 2 }}>
            <Box sx={{ display: { xs: 'none', md: 'flex' }, gap: 1 }}>
              {!isAuthenticated && (
                <>
                  <Button
                    variant="outlined"
                    size="sm"
                    component="a"
                    href="/login"
                    sx={{ color: 'white', borderColor: '#555', '&:hover': { backgroundColor: '#232f3e' } }}
                  >
                    Login
                  </Button>
                  <Button
                    variant="solid"
                    size="sm"
                    component="a"
                    href="/register"
                    sx={{ backgroundColor: '#febd69', color: '#111', '&:hover': { backgroundColor: '#f3a847' } }}
                  >
                    Register
                  </Button>
                </>
              )}
              <IconButton
                component="a"
                href="/cart"
                sx={{ color: 'white', '&:hover': { backgroundColor: 'orange', color: 'white' }  }}
              >
                <ShoppingCartCheckoutIcon sx={{fontSize:'2rem'}} />
              </IconButton>
              {isAuthenticated && (
                <Stack direction="row" spacing={2}>
                  <IconButton onClick={handleUserMenuOpen} sx={{ p: 0 }}>
                    <Avatar alt="User Avatar" src="/static/images/avatar/1.jpg" />
                  </IconButton>
                  <UserMenu
                    anchorEl={userMenuAnchorEl}
                    open={Boolean(userMenuAnchorEl)}
                    onClose={handleUserMenuClose}
                    onLogout={handleLogout}
                  />
                </Stack>
              )}
            </Box>       
        </Box>
      </Box>
      </Box>
      {/* Menu items moved here */}
      <Box sx={{ display:'flex',  gap: 1, justifyContent: 'center', flexGrow: 0, marginLeft: 0, backgroundColor: '#2a3444', paddingTop: 1, paddingBottom: 1, paddingLeft: 2, paddingRight: 2 }}>
        <Box>
          <IconButton
            variant="outlined"
            color="neutral"
            onClick={toggleAllMenuDrawer(true)}
            sx={{ color: 'white', borderColor: '#555', mr: 1}} // Adjusted for dark background and added margin
          >
          <MenuIcon sx={{ fontSize: '1.5rem' }}  />
          </IconButton>
        </Box>
        <Box>
        {menuItems.map((item) => (
          <Link
            key={item.label}
            href={item.href}
            level="body-sm"
            underline="hover"
            sx={{ color: 'white', px: 1, py: 0.5, '&:hover': { textDecoration: 'underline', backgroundColor: '#232f3e' }, fontSize: '1.1rem', marginLeft:0}} // Adjusted font size for better readability
          >
            {item.label}
          </Link>
        ))}
        </Box>
      </Box>
    </Box>       
    </Sheet>
  );
};