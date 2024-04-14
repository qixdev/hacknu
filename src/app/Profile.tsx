import { useState } from 'react';
import { Sidebar } from "../components/Sidebar/Sidebar.tsx";
import GoalCard from "../components/GoalCard.tsx";
import { Navbar } from "../components/Navbar.tsx";
import userIcon from '../../public/images/userIcon.webp';

const Profile = () => {
    const [showData, setShowData] = useState(false);

    const toggleData = () => {
        setShowData(!showData);
    };



    const response = {
        prompt: "Текст (от лат. textus — ткань; сплетение, сочетание) — зафиксированная на каком-либо материальном носителе человеческая мысль; в общем плане связная и полная последовательность символов.\n" +
            "\n" +
            "Существуют две основные трактовки понятия «текст»: имманентная (расширенная, философски нагруженная) и репрезентативная (более частная). Имманентный подход подразумевает отношение к тексту как к автономной реальности, нацеленность на выявление его внутренней структуры",
        files: ["../../../public/images/graphic.jpg"]
    };

    // const handleRequest = async (request) => {
    //     // Ваша логика обработки запроса здесь
    //
    //     // Возвращаем ответ
    //     return response;
    // };
    //
    // handleRequest(handleRequest)
    //     .then((result) => {
    //         console.log(result);
    //     })
    //     .catch((error) => {
    //         console.error("Error handling request:", error);
    //     });


    return (
        <div className="">
            {/*<Navbar/>*/}
            <Sidebar/>
            <div className="w-full flex flex-col justify-center items-center">
                <div className="relative w-[70%] h-48 mt-5 m-auto p-10 flex justify-between items-center rounded-lg bg-[#55bbeb] text-white">
                    <p className="absolute top-10 py-2 left-10 mt-[-1rem]">MY PROFILE</p>
                    <button className="absolute top-10 right-10 mt-[-1rem] px-5 py-2 rounded-lg bg-[#a7def9]" onClick={toggleData}>MY DATA</button>
                </div>
                <div className={`w-[50%] pb-20 flex flex-col justify-start items-center rounded-xl mt-[-100px] mb-10 bg-[#fff] ${showData ? 'opacity-100' : 'opacity-0'} transition-opacity duration-300`} style={{ zIndex: 5, boxShadow: '0 2px 2px 1px #a1a8b0' }}>
                    <div className="w-[80%] m-auto my-16 flex justify-between">
                        <div className={`w-[30%] flex flex-col justify-center items-center gap-4`}>
                            <img className={`w-[]`} src={userIcon}/>
                            <button className={`px-5 py-3 bg-gray-200 rounded-full border-gray-500`}>Upload photo</button>
                        </div>

                        <div className={`w-[60%] flex flex-col justify-between border-2 p-5 rounded-lg`}>
                            <div>
                                <label className={``}>Youe name</label>
                                <div className={`flex justify-between items-center`}>
                                    <p className={`font-bold`}>Sid</p>
                                    <button className={`px-10 py-2 bg-gray-200 rounded-full border-gray-500`}>Edit</button>
                                </div>
                            </div>
                            <div>
                                <label>Email</label>
                                <div className={`flex justify-between items-center`}>
                                    <p className={`font-bold`}>aba@mail.ru</p>
                                    <button className={`px-10 py-2 bg-gray-200 rounded-full border-gray-500`}>Edit</button>
                                </div>
                            </div>
                            <div>
                                <label>Phone Number</label>
                                <div className={`flex justify-between items-center`}>
                                    <p className={`font-bold`}>+ 7 778 432 12 21</p>
                                    <button className={`px-10 py-2 bg-gray-200 rounded-full border-gray-500`}>Edit</button>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div className={`w-[80%] pb-14 pt-5 mt-15 rounded-lg m-auto bg-[#fff] ${showData ? 'opacity-100' : 'opacity-0'} transition-opacity duration-300`} style={{ zIndex: 5, boxShadow: '1px 2px 2px 1px #a1a8b0' }}>
                        <h2 className={`text-[30px] font-bold text-center`}>Statistics</h2>
                        {
                            response.files.map(item => (
                                <img className={`w-[80%] m-auto mt-6`} src={item}/>
                            ))
                        }
                    </div>
                    <div className={`w-[80%] pb-14 pt-5 mt-20 rounded-lg m-auto bg-[#fff] ${showData ? 'opacity-100' : 'opacity-0'} transition-opacity duration-300`} style={{ zIndex: 5, boxShadow: '1px 2px 2px 1px #a1a8b0' }}>
                        <h2 className={`text-[30px] font-bold text-center`}>Recommendations</h2>
                        <p className={`mx-10 mt-5`}>{response.prompt}</p>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default Profile;


