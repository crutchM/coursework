import {useLocation} from "react-router-dom";
import {Box, Container, Typography, Link as MuiLink} from "@mui/material";
import fetchGithubUrl from "../utils/fetchGithubUrl";


const LoginPage = () =>{
    //const location = useLocation()
    let from =  '/'

    return (
        <Container
            maxWidth={false}
            sx={{
                display: 'flex',
                justifyContent: 'center',
                alignItems: 'center',
                height: '100vh',
                backgroundColor: '#2363eb',
            }}
        >
            <Box width='27rem'>
                <Typography
                    variant='h6'
                    component='p'
                    sx={{
                        my: '1.5rem',
                        textAlign: 'center',
                        color: 'white',
                    }}
                >
                    Log in with another provider:
                </Typography>
                <Box
                    width='100%'
                    sx={{
                        backgroundColor: '#e5e7eb',
                        p: { xs: '1rem', sm: '2rem' },
                        borderRadius: 2,
                    }}
                >
                    <MuiLink
                        href={fetchGithubUrl(from)}
                        sx={{
                            backgroundColor: '#f5f6f7',
                            borderRadius: 1,
                            py: '0.6rem',
                            mt: '1.5rem',
                            columnGap: '1rem',
                            textDecoration: 'none',
                            color: '#393e45',
                            cursor: 'pointer',
                            fontWeight: 500,
                            '&:hover': {
                                backgroundColor: '#fff',
                                boxShadow: '0 1px 13px 0 rgb(0 0 0 / 15%)',
                            },
                        }}
                        display='flex'
                        justifyContent='center'
                        alignItems='center'
                    >
                        {/*<GitHubLogo style={{ height: '2rem' }} />*/}
                        GitHub
                    </MuiLink>
                </Box>
            </Box>

        </Container>
    )
}


export default LoginPage