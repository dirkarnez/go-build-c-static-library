cmake_minimum_required(VERSION 3.14)

project(sumcpp)

set(CMAKE_CXX_STANDARD 11)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

add_library(sum STATIC IMPORTED)

set_property(
  TARGET sum APPEND PROPERTY IMPORTED_CONFIGURATIONS DEBUG
)

set_target_properties(sum PROPERTIES
  IMPORTED_LINK_INTERFACE_LANGUAGES_DEBUG "C"
  IMPORTED_LOCATION_DEBUG "${SUM_ROOT}/sum.a"
)

set_property(
  TARGET sum APPEND PROPERTY IMPORTED_CONFIGURATIONS RELEASE
)

set_target_properties(sum PROPERTIES
  IMPORTED_LINK_INTERFACE_LANGUAGES_RELEASE "C"
  IMPORTED_LOCATION_RELEASE "${SUM_ROOT}/sum.a"
)
  
set_target_properties(sum PROPERTIES
  MAP_IMPORTED_CONFIG_MINSIZEREL Release
  MAP_IMPORTED_CONFIG_RELWITHDEBINFO Release
)

target_include_directories(sum INTERFACE "${SUM_ROOT}")

add_executable(${PROJECT_NAME} 
    "src/main.cpp"
)

target_link_libraries(${PROJECT_NAME} sum)