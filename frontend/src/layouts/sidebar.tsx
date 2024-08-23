import React from 'react';
import Dropdown from '../components/dropdown';

const Sidebar: React.FC = () => {
  const categoriesDropdownItems = [
    { label: 'JEANS', href: '#' },
    { label: 'T-SHIRTS', href: '#' },
    { label: 'SWIMWEAR', href: '#' },
    { label: 'SHIRTS', href: '#' },
    { label: 'SHORTS', href: '#' },
    { label: 'PANTS', href: '#' },
    { label: 'GRAPHIC T-SHIRTS', href: '#' },
    { label: 'ACTIVEWEAR', href: '#' },
  ];

  const priceDropdownItems = [
    { label: '- C$20.00', href: '#' },
    { label: 'C$20.00 - C$40.00', href: '#' },
    { label: 'C$40.00 - C$70.00', href: '#' },
    { label: 'C$70.00 - C$100.00', href: '#' },
    { label: 'C$100.00 - C$300.00', href: '#' },
    { label: 'C$300.00 -', href: '#' }
  ];

  const colorDropdownItems = [
    { label: 'Beige', href: '#' },
    { label: 'Black', href: '#' },
    { label: 'Blue', href: '#' },
    { label: 'Brown', href: '#' },
    { label: 'Cyan', href: '#' },
    { label: 'Dark Blue', href: '#' },
    { label: 'Dark Green', href: '#' },
    { label: 'Dark Grey', href: '#' },
    { label: 'Gold', href: '#' },
    { label: 'Green', href: '#' },
    { label: 'Grey', href: '#' },
    { label: 'Indigo', href: '#' },
    { label: 'Ivory', href: '#' },
    { label: 'Maroon', href: '#' },
    { label: 'Navy', href: '#' },
    { label: 'Olive', href: '#' },
    { label: 'Orange', href: '#' },
    { label: 'Pink', href: '#' },
    { label: 'Purple', href: '#' },
    { label: 'Red', href: '#' },
    { label: 'White', href: '#' }
  ];
  return (
    <div className="sidebar">
      <h2 style={{ paddingBottom: '15px' }} >220 items</h2>
      <ul className="navbar-menu">
        <li>
          <Dropdown label="Categories" items={categoriesDropdownItems} />
        </li>
        <li>
          <Dropdown label="Price" items={priceDropdownItems} />
        </li>
        <li>
          <Dropdown label="Color" items={colorDropdownItems} />
        </li>
      </ul>
    </div>
  );
};

export default Sidebar;
