const Navbar = (props: any) => {
    return (
        <nav className="navbar">
          <h1>{props.title}</h1>
          <p>{props.paragraf}</p>
        </nav>
      );
}
 
export default Navbar;