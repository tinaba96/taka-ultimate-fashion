import React, { useState, useRef, useEffect, MouseEvent } from 'react';

interface DropdownItem {
  label: string;
  href: string;
}

interface DropdownProps {
  label: string;
  items: DropdownItem[];
}

const Dropdown: React.FC<DropdownProps> = ({ label, items }) => {
  const [isOpen, setIsOpen] = useState<boolean>(false);
  const dropdownRef = useRef<HTMLDivElement | null>(null);

  const toggleDropdown = () => {
    setIsOpen(!isOpen);
  };

// function handleClickOutside(this: Document, event: MouseEvent): void {
//     if (dropdownRef.current && !dropdownRef.current.contains(event.target as Node)) {
//       setIsOpen(false);
//     }
//   };

//   useEffect(() => {
//       document.removeEventListener('mousedown', handleClickOutside as EventListener);
//   }, []);
function handleClickOutside(this: Document, event: MouseEvent): void {
  if (dropdownRef.current && !dropdownRef.current.contains(event.target as Node)) {
    setIsOpen(false);
  }
}
  useEffect(() => {
  document.addEventListener('mousedown', handleClickOutside as unknown as EventListener);
  return () => {
    document.removeEventListener('mousedown', handleClickOutside as unknown as EventListener);
  };
}, []);

  return (
    <div className="dropdown" ref={dropdownRef}>
      <button className="dropdown-button" onClick={toggleDropdown}>
        {label}
      </button>
      {isOpen && (
        <div className="dropdown-menu">
          {items.map((item, index) => (
            <a key={index} href={item.href}>{item.label}</a>
          ))}
        </div>
      )}
    </div>
  );
};

export default Dropdown;
