import React, {useEffect, useState} from 'react';


// const products = [
//     {
//       id: 1,
//       name: 'Basic Tee',
//       href: '#',
//       imageSrc: 'https://tailwindui.com/img/ecommerce-images/product-page-01-related-product-01.jpg',
//       imageAlt: "Front of men's Basic Tee in black.",
//       price: '$35',
//       color: 'Black',
//     },
//     // More products...
//   ]

  interface Product {
    ID: number;
    Name: string;
    // href: string;
    ImageURL: string;
    // imageAlt: string;
    Price: string;
    // color: string;
  }

  interface DataResponse {
    datas: Product[];
}

const Home: React.FC = () => {
    const [products, setProducts] = useState<Product[]>([])
    // const [products2, setProducts2] = useState<Product[]>([
    //     {
    //         id: 123,
    //         name: '',
    //         href: '',
    //         imageSrc: '',
    //         imageAlt: '',
    //         price: '',
    //         color: '',
    //       },
    // ])

    useEffect(() => {
        fetch('http://localhost:8888/products')
        .then(response => response.json())
        .then((data: DataResponse) => {
            setProducts(data.datas)
            // console.log(data)
        })
        .catch(error => {
            console.error('Request Error(Please check the response type):', error);
        });
    }, [])


  return (
    <div>
      <h1 className="text-2xl header-title">T-SHIRTS</h1>
    <div className="bg-white">
      <div className="mx-auto max-w-2xl px-4 py-16 sm:px-6 sm:py-24 lg:max-w-7xl lg:px-8">
        <h2 className="text-2xl font-bold tracking-tight text-gray-900">Customers also purchased</h2>
        <div className="mt-6 grid grid-cols-1 gap-x-6 gap-y-10 sm:grid-cols-2 lg:grid-cols-4 xl:gap-x-8">
          {products.map((product) => (
              <div key={product.ID} className="group relative">
              <div className="aspect-h-1 aspect-w-1 w-full overflow-hidden rounded-md bg-gray-200 lg:aspect-none group-hover:opacity-75 lg:h-80">
                <img
                  alt={product.Name}
                  src={`https://oldnavy.gapcanada.ca/${product.ImageURL}`}
                  className="h-full w-full object-cover object-center lg:h-full lg:w-full"
                />
              </div>
              <div className="mt-4 flex justify-between">
                <div>
                  <h3 className="text-sm text-gray-700">
                    <a href={product.ImageURL}>
                      <span aria-hidden="true" className="absolute inset-0" />
                      {product.Name}
                    </a>
                  </h3>
                  <p className="mt-1 text-sm text-gray-500">{product.Name}</p>
                </div>
                <p className="text-sm font-medium text-gray-900">{product.Price}</p>
              </div>
            </div>
          ))}
        </div>
      </div>
          </div>
          </div>
  );
};

export default Home;
