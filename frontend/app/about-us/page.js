"use client"
import React, { useState, useEffect } from 'react';
import Navbar from '../../components/navbar.js';

export default function AboutUs() {
    const members = [
        {
            name: 'Aland Mulia Pratama',
            nim: '13522124',
            role: 'Frontend Developer',
            email: '13522124@std.stei.itb.ac.id',
            image: '/aland.jpg',
        },
        {
            name: 'Rizqika Mulia Pratama',
            nim: '13522126',
            role: 'Simulated Annealing & Backend Developer',
            email: '13522126@std.stei.itb.ac.id',
            image: '/qika.jpg',
        },
        {
          name: 'Christian Justin Hendrawan',
          nim: '13522135',
          role: 'Genetic Algorithm',
          email: '13522150@std.stei.itb.ac.id',
          image: '/chris.jpg',
      },
        {
            name: 'Ikhwan Al Hakim',
            nim: '13522147',
            role: 'Hill Climbing Algorithm',
            email: '13522147@std.stei.itb.ac.id',
            image: '/ikhwan.jpg',
        },
    ];

    const [activeIndex, setActiveIndex] = useState(0);

    useEffect(() => {
        const interval = setInterval(() => {
            setActiveIndex((prevIndex) => (prevIndex + 1) % members.length);
        }, 3000); // Change every 3 seconds

        return () => clearInterval(interval);
    }, [members.length]);

    const getCardPosition = (index) => {
        const pos = (index - activeIndex + members.length) % members.length;

        switch (pos) {
            case 0:
                return 'translate-x-[-150%] scale-75 opacity-50 z-0'; // Left side
            case 1:
                return 'translate-x-0 scale-100 opacity-100 z-10'; // Center (front)
            case 2:
                return 'translate-x-[150%] scale-75 opacity-50 z-0'; // Right side
            default:
                return 'hidden'; // Hide other cards
        }
    };

    return (
        <div className="bg-gray-100 text-gray-800 font-sans min-h-screen pb-12">
            <Navbar />
            <div className="text-center mt-10">
                <h1 className="text-5xl font-bold text-gray-800 underline mb-4">LEMANSPEDIA CONTRIBUTORS</h1>
                <h3 className="text-xl text-gray-700 mx-[50px] leading-relaxed">
                    The main objective of this project is to create a WikiRace program with Breadth First Search (BFS) and Iterative Deepening Search (IDS) algorithms.
                </h3>
            </div>

            {/* Carousel Container */}
            <div className="relative w-full flex justify-center items-center mt-12 h-[400px] overflow-hidden">
                {members.map((member, index) => (
                    <div
                        key={index}
                        className={`absolute text-center transition-all duration-500 ease-in-out ${getCardPosition(index)} transform`}
                        style={{
                            transition: 'transform 0.5s ease, opacity 0.5s ease',
                        }}
                    >
                        <div className="card bg-white border border-gray-300 rounded-lg shadow-lg w-[300px] h-[400px] p-6 mx-4">
                            <div className="flex flex-col items-center">
                                <img src={member.image} className="h-[150px] w-[150px] rounded-full object-cover mb-6" alt={`${member.name}`} />
                                <p className="text-2xl font-bold mb-2">{member.name}</p>
                                <p className="text-gray-600">{member.nim}</p>
                                <p className="italic text-gray-700">{member.role}</p>
                                <p className="text-gray-600">({member.email})</p>
                            </div>
                        </div>
                    </div>
                ))}
            </div>
        </div>
    );
}
