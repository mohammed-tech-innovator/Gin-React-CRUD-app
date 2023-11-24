import React from "react";
import Logo from "../assets/Images/Logo.jpg"

const Navbar = ()=>{
    return(
        <nav className="navbar navbar-expand-lg navbar-light bg-light">

                <div className="collapse navbar-collapse" id="navbarTogglerDemo01">
                    <a className="navbar-brand mx-4" href="/">
                        <img className="mx-2 rounded-3" src = {Logo} style={{width:"40px"}} alt="Logo"/>
                        Hidden brand
                    </a>
                    <MyNavbarLinks/>
                </div>
        </nav>
    )
}


const MyNavbarLinks = () => {
    return (
      <div className="text-center">
        <ul className="navbar-nav mr-auto mt-2 mt-lg-0">
          <NavItem active={true} link="/">Home</NavItem>
          <NavItem link="/">Link</NavItem>
          <NavItem link="/" disabled={true}>Disabled</NavItem>
        </ul>
      </div>
    );
  }


const NavItem = ({ link, active, disabled, children }) => {
    const navItemClass = `nav-item${active ? ' active' : ''}`;
    const navLinkClass = `nav-link${disabled ? ' disabled' : ''}`;
  
    return (
      <li className={navItemClass}>
        <a className={navLinkClass} href={link}>
          {children}
          {active && <span className="sr-only">(current)</span>}
        </a>
      </li>
    );
  };


  export default Navbar;