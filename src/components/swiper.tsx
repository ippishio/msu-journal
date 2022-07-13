import React, { useRef, useState } from "react";
// Import Swiper React components
import { Swiper, SwiperSlide } from "swiper/react";

// Import Swiper styles
import "swiper/css";
import "swiper/css/pagination";
import "swiper/css/navigation";

import '../img/1.jpg'
// import required modules
import { Pagination, Navigation, Autoplay } from "swiper";
import { render } from "@testing-library/react";

export default class Sw extends React.Component{
    render(){
  return (
    <>
      <Swiper
      spaceBetween={30}
      centeredSlides={true}
      autoplay={{
        delay: 2500,
        disableOnInteraction: false,
      }}
        pagination={true
        }
        navigation={true}
        modules={[Pagination, Navigation,Autoplay]}
        className="mySwiper"
      >
        <SwiperSlide>
            <img src='1.jpg'></img>
           
           </SwiperSlide>
        
        
      </Swiper>
    </>
  );
    }
}