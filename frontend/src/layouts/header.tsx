import React from 'react';
import { Link } from 'react-router-dom';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faGlobe, faUser, faGem, faHeart, faShoppingBag } from '@fortawesome/free-solid-svg-icons';


const Header: React.FC = () => {
  return (
    <header>
        <div className='header-comp py-6'>
          <div className='top-header-comp py-6'>
          <nav className="nav">
          <ul className="nav-list">
            <li className="nav-item">
              <Link to="/" className="nav-link">WOMEN</Link>
            </li>
            <li className="nav-item">
              <Link to="/about" className="nav-link">MEN</Link>
            </li>
            <li className="nav-item">
              <Link to="/services" className="nav-link">KIDS</Link>
            </li>
            <li className="nav-item">
              <Link to="/contact" className="nav-link">HOME</Link>
            </li>
          </ul>
        </nav>
        <nav className="nav">
          <ul className="nav-list">
            <li className="nav-item">
            <Link to="/" className="nav-link">
              <FontAwesomeIcon icon={faGlobe} />
            </Link>
            </li>
            <li className="nav-item">
            <Link to="/" className="nav-link">
              <FontAwesomeIcon icon={faUser} />
            </Link>
            </li>
            <li className="nav-item">
            <Link to="/" className="nav-link">
              <FontAwesomeIcon icon={faGem} />
            </Link>
            </li>
            <li className="nav-item">
            <Link to="/" className="nav-link">
              <FontAwesomeIcon icon={faHeart} />
            </Link>
            </li>
            <li className="nav-item">
            <Link to="/" className="nav-link">
              <FontAwesomeIcon icon={faShoppingBag} />
            </Link>
            </li>
          </ul>
        </nav>
              <h1 className="header-comp-title font-sans text-2xl font-bold tracking-tight text-gray-900">Taka's Ultimate Fashion</h1>
              </div><div>
          <nav className="nav">
          <ul className="nav-list">
            <li className="nav-item">
              <Link to="/" className="nav-link">WHATS NEW</Link>
            </li>
            <li className="nav-item">
              <Link to="/about" className="nav-link">BRANDS</Link>
            </li>
            <li className="nav-item">
              <Link to="/services" className="nav-link">CLOTHING</Link>
            </li>
            <li className="nav-item">
              <Link to="/contact" className="nav-link">BAGS</Link>
            </li>
            <li className="nav-item">
              <Link to="/contact" className="nav-link">SHOES</Link>
            </li>
            <li className="nav-item">
              <Link to="/contact" className="nav-link">ACCESSORIES</Link>
            </li>
            <li className="nav-item">
              <Link to="/contact" className="nav-link">SPORTS</Link>
            </li>
            <li className="nav-item">
              <Link to="/contact" className="nav-link"><span style={{ color: 'red' }}>SALES</span></Link>
            </li>
          </ul>
        </nav>
          </div>
        </div>
    </header>
  );
};

export default Header;
