import * as React from 'react';
import Box from '@mui/material/Box';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import ListItemButton from '@mui/material/ListItemButton';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import { ListItemAvatar } from '@mui/material';
import Divider from '@mui/material/Divider';
import InboxIcon from '@mui/icons-material/Inbox';
import DraftsIcon from '@mui/icons-material/Drafts';
import { AccountCircle } from '@mui/icons-material';
import { Avatar } from '@mui/material';

export default function BasicList() {
  return (
    <Box className='my-5 px-5 drop-shadow-sm'>
        <List>
          <ListItem >
              <ListItemAvatar>
                  <Avatar className='bg-red-500'>
                <AccountCircle />
                  </Avatar>
              </ListItemAvatar>
              <ListItemText primary="Pasit Sri-intarasut" />
          </ListItem>
      <Divider />
          <ListItem disablePadding>
            <ListItemButton>
              <ListItemIcon>
                <InboxIcon />
              </ListItemIcon>
              <ListItemText primary="Table" />
            </ListItemButton>
          </ListItem>
          <ListItem disablePadding>
            <ListItemButton>
              <ListItemIcon>
                <DraftsIcon />
              </ListItemIcon>
              <ListItemText primary="Drafts" />
            </ListItemButton>
          </ListItem>
        </List>
    </Box>
  );
}