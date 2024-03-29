import React, { useEffect, useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import Main from "../components/Main";
import { UserAuth } from "../context/AuthContext";

const Signup = ({ viewValue, toggleNavbar, toggleLoggedIn }) => {
  const [navbar, setNavbar] = useState("NavbarSignUp");
  const navigate = useNavigate();

  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const { user, signUp } = UserAuth();
  const [error, setError] = useState("");

  const handleSignup = async (e) => {
    e.preventDefault();
    try {
      await signUp(email, password);
      toggleNavbar("NavbarHome");
      toggleLoggedIn(true);
      navigate("/");
    } catch (error) {
      console.log(error);
      setError(error);
    }
  };

  useEffect(() => {
    // Check if we are on the /display route
    if (window.location.pathname === "/login") {
      // Set the navbar to 'NavbarDisplay' if on the /display route
      setNavbar("NavbarLogIn");
    } else if (window.location.pathname === "/") {
      setNavbar("NavbarHome");
    } else if (window.location.pathname === "/display") {
      setNavbar("NavbarDisplay");
    } else if (window.location.pathname === "/account") {
      setNavbar("NavbarAccount");
    } else if (window.location.pathname === "/signup") {
      setNavbar("NavbarSignUp");
    }
  }, []);

  useEffect(() => {
    // Check if we should navigate to /display
    if (navbar === "NavbarHome") {
      navigate("/");
    } else if (navbar === "NavbarDislay") {
      navigate("/display");
    } else if (navbar === "NavbarAccount") {
      navigate("/account");
    } else if (navbar === "NavbarLogIn") {
      navigate("/login");
    } else if (navbar === "NavbarSignUp") {
      navigate("/signup");
    }
  }, [navbar, navigate]);
  return (
    <>
      <Main />
      <div className="mt-[-500px] w-full h-screen">
        <div className="absolute"></div>
        <div className="relative w-full px-4 py-24 z-50">
          <div className="max-w-[450px] h-[600px] mx-auto bg-gray-400 text-black-600 rounded-lg">
            <div className="max-w-[320px] mx-auto py-16">
              <h1 className="text-3xl font-bold">Register</h1>
              {error ? <p className="p-3 bg-blue-400">{error}</p> : null}
              <form
                onSubmit={handleSignup}
                className="w-full flex flex-col py-4"
              >
                <input
                  onChange={(e) => setEmail(e.target.value)}
                  className="p-3 my-2 bg-gray-600 rounded"
                  type="email"
                  placeholder="Email"
                  autoComplete="email"
                />
                <input
                  onChange={(e) => setPassword(e.target.value)}
                  className="p-3 my-2 bg-gray-600 rounded"
                  type="password"
                  placeholder="Password"
                  autoComplete="current-password"
                />

                <button className="bg-blue-600 py-3 my-6 rounded font-bold ">
                  Sign Up
                </button>

                <p className="py-4">
                  <span className="text-black-600">Already registered?</span>
                  {"  "}

                  <Link to="/login">
                    <button
                      onClick={() => {
                        setNavbar("NavbarLogIn");
                      }}
                    >
                      Sign In
                    </button>
                  </Link>
                </p>
              </form>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default Signup;

/*import React, { useEffect, useState } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import Main from '../components/Main'
import {UserAuth} from '../context/AuthContext'

const Signup = () => {
    const [email,setEmail] = useState('')
    const [password,setPassword] =useState('') 
    const [user,signUp] = UserAuth()

    const handleSubmit=async(e)=>{
        e.preventDefault()
        try{
            await signUp(email,password)
        }catch(error){
            console.log(error)
        }
    }

    const  [navbar,setNavbar] = useState('NavbarSignUp');
    const navigate = useNavigate();



      useEffect(() => {
      // Check if we are on the /display route
      if (window.location.pathname === '/login') {
            // Set the navbar to 'NavbarDisplay' if on the /display route
            setNavbar('NavbarLogIn');
      }else if(window.location.pathname === '/'){
            setNavbar('NavbarHome');
      }else if(window.location.pathname === '/display'){
        setNavbar('NavbarDisplay');
        }else if(window.location.pathname === '/account'){
        setNavbar('NavbarAccount');
        }else if(window.location.pathname === '/signup'){
            setNavbar('NavbarSignUp');
        }
      }, []);
      
      useEffect(() => {
      // Check if we should navigate to /display
      if (navbar === 'NavbarHome') {
            navigate('/');
      }else if(navbar === 'NavbarDislay'){
            navigate('/display');
      }else if(navbar === 'NavbarAccount'){
            navigate('/account');
      }else if(navbar === 'NavbarLogIn'){
        navigate('/login');
        }
        else if(navbar === 'NavbarSignUp'){
            navigate('/signup');
            }
      }, [navbar, navigate]);
  return (
    <>
        <Main/>
        <div className = "mt-[-500px] w-full h-screen">

            <div className='absolute'></div>
            <div className = 'relative w-full px-4 py-24 z-50'>
                <div className= 'max-w-[450px] h-[600px] mx-auto bg-gray-400 text-black-600 rounded-lg'>
                    <div className ='max-w-[320px] mx-auto py-16'>
                        <h1 className='text-3xl font-bold'>Register</h1>
                        <form onSubmit={handleSubmit} className='w-full flex flex-col py-4'>
                            <input onChange ={(e)=>setEmail(e.target.value)} className = 'p-3 my-2 bg-gray-600 rounded' type = 'email' placeholder = 'Email' autoComplete='email' />
                            <input onChange ={(e)=>setPassword(e.target.value)} className = 'p-3 my-2 bg-gray-600 rounded' type='password' placeholder='Password'autoComplete='current-password'/>

                            <button className='bg-blue-600 py-3 my-6 rounded font-bold '>Login</button>

                            <p className = 'py-4'><span className='text-black-600'>Already registered?</span>{'  '}
                            
                                <Link to = '/login'>Sign in</Link>
                            </p>

                        </form>

                    </div>

                </div>

            </div>


        </div>
    </>

  )
}

export default Signup*/
