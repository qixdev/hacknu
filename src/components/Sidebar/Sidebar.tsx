import {useState} from "react";
import './Sidebar.css';
import logo from '../../../public/images/beyimicon.svg';
import burger from '../../../public/images/menu.png'


const navItems = [
    "home",
    "data"
]



export const Sidebar = () => {
    const [isOpen, setIsOpen] = useState(false);

    return (
        <aside className={`sidebar ${isOpen ? "open" : ""}`}>
            <div className={"sidebar-inner"}>
                <header className={"sidebar-header"}>
                    <button
                        type={"button"}
                        className={"sidebar-burger"}
                        onClick={() => setIsOpen(!isOpen)}
                    >
                        {/*<img className={'sidebar-burger-icon'} src={burger} width={20} height={20}/>*/}
                        {isOpen ? (<img src={logo} className={"sidebar-logo"}/>) : (<img className={'sidebar-burger-icon'} src={burger} width={20} height={20}/>)}
                    </button>
                    {/*<img src={logo} className={"sidebar-logo"}/>*/}
                </header>
                <nav className={"sidebar-menu"}>
                    {navItems.map(item => (
                        <button className={"sidebar-button"}>
                            <p>{item}</p>
                        </button>
                    ))}
                </nav>
            </div>
        </aside>
    );
}