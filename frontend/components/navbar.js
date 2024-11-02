import Image from 'next/image';
import Link from 'next/link';
import icon from '../public/Lemanspedia-removebg.png';

const Navbar = () => {
    return (
        <nav className="sticky top-0 bg-white bg-opacity-90 shadow-md z-20 flex items-center justify-between px-6 py-4">
            {/* Logo Section */}
            <div className="flex items-center">
                <Image src={icon} alt="Lemanspedia Logo" width={60} height={60} />
                <h1 className="ml-4 font-bold text-2xl text-gray-800">Magic Cube Solver</h1>
            </div>
            
            {/* Links Section */}
            <ul className="flex space-x-8 text-gray-800">
                <li>
                    <Link href="/" className="hover:text-gray-500 font-bold text-lg transition-all duration-200">
                        Home
                    </Link>
                </li>
                <li>
                    <Link href="/magic-cube" className="hover:text-gray-500 font-bold text-lg transition-all duration-200">
                        Magic Cube
                    </Link>
                </li>
                <li>
                    <Link href="/about-us" className="hover:text-gray-500 font-bold text-lg transition-all duration-200">
                        About Us
                    </Link>
                </li>
            </ul>
        </nav>
    );
}

export default Navbar;