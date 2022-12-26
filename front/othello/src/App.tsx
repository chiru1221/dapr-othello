import { useState } from 'react';
import Modal from '@mui/material/Modal';
import Box from '@mui/material/Box';
import Setting from './setting';
import Main from './main';
import './App.css';

const App = () => {
  const [setting, setSetting] = useState<boolean>(true)
  const [firstAttack, setFirstAttack] = useState<boolean>(true)
  return (
    <div>
      <Modal
        open={setting}
      >
        <Box sx={{
          width: '50vw',
          height: '50vh',
          position: 'absolute',
          top: '50%',
          left: '50%',
          transform: 'translate(-50%, -50%)',
          backgroundColor: 'white',
          border: '2px solid #000',
          boxShadow: 24,
        }}>
          <Setting
            setSetting={setSetting}
            setFirstAttack={setFirstAttack}
          />
        </Box>
      </Modal>
      {!setting && <Main firstAttack={firstAttack} />}
    </div>
  );
}

export default App;
