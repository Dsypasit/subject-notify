import * as React from "react";
import Accordion from "@mui/material/Accordion";
import AccordionSummary from "@mui/material/AccordionSummary";
import AccordionDetails from "@mui/material/AccordionDetails";
import Typography from "@mui/material/Typography";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import {
  Button,
  TextField,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  Box,
} from "@mui/material";

export default function Command() {
  const [day, setDay] = React.useState("");
  const [time, setTime] = React.useState("");

  const handleChangeDay = (event) => {
    setDay(event.target.value);
  };
  const handleChangeTime = (event) => {
    setTime(event.target.value);
  };
  return (
    <div className="px-10">
      <Accordion>
        <AccordionSummary
          expandIcon={<ExpandMoreIcon />}
          aria-controls="panel1a-content"
          id="panel1a-header"
        >
          <Typography>Create subject</Typography>
        </AccordionSummary>
        <AccordionDetails>
          <div className="grid">
            <div className="flex justify-around">
              <div className="mx-5 flex flex-wrap items-center">
                <Typography className="mr-3">Subject name:</Typography>
                <TextField
                  required
                  id="standard-basic"
                  label="subject name"
                  variant="standard"
                />
              </div>
              <div className="mx-5 flex items-center">
                <Typography className="mr-3">Day:</Typography>
                <Box sx={{ minWidth: 120 }}>
                  <FormControl fullWidth>
                    <InputLabel required id="demo-simple-select-label">Day</InputLabel>
                    <Select
                      required
                      labelId="demo-simple-select-label"
                      id="demo-simple-select"
                      value={day}
                      label="Day"
                      onChange={handleChangeDay}
                    >
                      <MenuItem value={"Mon"}>Mon</MenuItem>
                      <MenuItem value={"Tue"}>Tue</MenuItem>
                      <MenuItem value={"Wen"}>Wen</MenuItem>
                      <MenuItem value={"Thu"}>Thu</MenuItem>
                      <MenuItem value={"Fri"}>Fri</MenuItem>
                    </Select>
                  </FormControl>
                </Box>
              </div>
              <div className="mx-5 flex items-center">
                <Typography className="mr-3">Time:</Typography>
                <Box sx={{ minWidth: 120 }}>
                  <FormControl fullWidth>
                    <InputLabel required id="demo-simple-select-label">Time</InputLabel>
                    <Select
                      required
                      labelId="demo-simple-select-label"
                      id="demo-simple-select"
                      value={time}
                      label="time"
                      onChange={handleChangeTime}
                    >
                      <MenuItem value={"Morn"}>Morn</MenuItem>
                      <MenuItem value={"After"}>Afer</MenuItem>
                    </Select>
                  </FormControl>
                </Box>
              </div>
              <div className="mx-5 flex items-center">
                <Typography className="mr-3">Link: </Typography>
                <TextField
                  id="standard-basic"
                  label="subject name"
                  variant="standard"
                />
              </div>
              <div className="mx-5 flex items-center">
                <Typography className="mr-3">Teacher: </Typography>
                <TextField
                  id="standard-basic"
                  label="subject name"
                  variant="standard"
                />
              </div>
            </div>
            <Button className="mt-5" variant="outlined">
              Submit
            </Button>
          </div>
        </AccordionDetails>
      </Accordion>
      <Accordion>
        <AccordionSummary
          expandIcon={<ExpandMoreIcon />}
          aria-controls="panel2a-content"
          id="panel2a-header"
        >
          <Typography>Accordion 2</Typography>
        </AccordionSummary>
        <AccordionDetails>
          <Typography>
            Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse
            malesuada lacus ex, sit amet blandit leo lobortis eget.
          </Typography>
        </AccordionDetails>
      </Accordion>
    </div>
  );
}
