import { Box, ButtonBase } from "@mui/material";
import AnimatedCircle from "../base/AnimatedCircle";

const KaspaLogo = () => {
    return (
        <AnimatedCircle>
            <ButtonBase color="primary" sx={{borderRadius: '50%'}} focusRipple>
                <Box sx={{
                    borderRadius: '50%',
                    borderStyle: 'solid',
                    borderColor: '#ffffff',
                    borderWidth: '6px',
                    height: '92px',
                    backgroundColor: '#ffffff'
                }}>
                    <svg xmlns="http://www.w3.org/2000/svg" width="80" height="80" viewBox="-4 -4 132 132">
                        <g>
                            <path d="M116.27,38.3A59.47,59.47,0,0,0,103.17,19c-5.4-5.4-12.4-9.5-19.6-12.5a58.81,58.81,0,0,0-22.7-4.3c-8,0-16.2.3-23.2,3.2-7.2,3-13.3,8.7-18.7,14.1S7.17,30.7,4.17,37.9c-2.8,7-1.9,15.5-1.9,23.5s.6,15.8,3.5,22.8c3,7.3,9.1,12.4,14.5,17.8s10.4,11.9,17.6,14.9a61.11,61.11,0,0,0,23,4.9c8,0,15.8-2.4,22.8-5.3a60.41,60.41,0,0,0,32-32.4c2.9-7,6.2-14.7,6.2-22.7S119.17,45.4,116.27,38.3ZM77.57,95.6l-12.7-1.9,3.7-24.6L42,89.6l-7.8-10.2,23.3-18L34.17,43.5,42,33.3l26.6,20.5-3.7-24.6,12.7-1.9,5.1,34.1Z" style={{fill:"#71c9bb"}}/>
                        </g>
                    </svg>
                </Box>
            </ButtonBase>
        </AnimatedCircle>
    );
}

export default KaspaLogo;
