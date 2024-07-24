import React, {useEffect, useState} from 'react';
import Select from "react-tailwindcss-select";

const options: Option[] = [
    { value: "JEANS", label: "JEANS" },
    { value: "T-SHIRTS", label: "T-SHIRTS" },
    { value: "Honeybee", label: "🐝 Honeybee" }
];
interface Option {
    value: string;
    label: string;
}

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
    // const [dropdownOpen, setDropdownOpen] = useState(false);
    const [categories, setCategories] = useState<Option[]>([]);
    // const [categories, setCategories] = useState<string[]>([]);

    // const handleChange = (value: String) => {
    //     console.log("value:", value);
    //     setAnimal(value);
    // };
    const handleChange = (selectedOption: Option[]) => {
      setCategories(selectedOption as Option[]);
      const selectedCategories = selectedOption.length !== 0 ? selectedOption.map(option => option.value) : []
      // setCategories(selectedCategories as string[])
      fetch(`http://localhost:8888/products?categories=${selectedCategories}`)
        .then(response => response.json())
        .then((data: DataResponse) => {
            setProducts(data.datas)
            // this.setState({ products: response.data });
            // console.log(data)
        })
        .catch(error => {
            console.error('Request Error(Please check the response type):', error);
        });
      }

      const handleClick = async (e: React.MouseEvent<HTMLImageElement>, product_id:number) => {
        e.preventDefault();
    
        const data = {
          ProductID: product_id,
          Status: 1
        };
        console.log('datta')
        console.log(data)
    
        // try {
        //   const res = await fetch('http://localhost:8080/saved', {
        //     method: 'POST',
        //     headers: {
        //       'Content-Type': 'application/json',
        //     },
        //     body: JSON.stringify(data),
        //   });
    
        //   await res.json();
        //   // const result = await res.json();
        //   // setResponse(result);
        // } catch (error) {
        //   console.error('Error:', error);
        // }
      };

    // const toggleDropdown = () => {
    //     setDropdownOpen(!dropdownOpen);
    // };
    // const headerText = categories.map(category => category.value).join(' & ');
    console.log('cate')
    console.log(categories)
    const headerText = (categories && categories?.length !== 0) ? categories.map(category => category.value).join(' & ') : "ALL";

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
        <div className="max-w-2xl mx-auto py-0 sm:px-6 sm:py-16 lg:max-w-7xl lg:px-8">
            <div className="flex items-center space-x-4">
                <span>SEARCH </span>
                <Select
                    primaryColor={"indigo"}
                    value={categories}
                    onChange={handleChange}
                    options={options}
                    isMultiple={true}
                />
                <button className="bg-green-700 hover:bg-green-800 text-white font-bold py-2 px-4 rounded">
                      🧠 AI 🤖 Recommend
                </button>
            </div>

        <h1 className="py-10 text-2xl header-title">{products.length !== 0 ? headerText : ''}</h1>
            <div  className="grid grid-cols-1 gap-x-6 gap-y-10 sm:grid-cols-2 lg:grid-cols-4 xl:gap-x-8">
            {products.length !== 0 ? (
              products.map((product) => (
                  <div key={product.ID} className="group relative">
                  <div className="aspect-h-1 aspect-w-1 w-full overflow-hidden rounded-md bg-gray-200 lg:aspect-none group-hover:opacity-75 lg:h-80">
                      <img
                      alt={product.Name}
                      src={`https://oldnavy.gapcanada.ca/${product.ImageURL}`}
                      className="h-full w-full object-cover object-center lg:h-full lg:w-full"
                      onClick={(e) => handleClick(e, product.ID)}
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
                      <p className="text-sm font-medium text-gray-900">{product.Price}</p>
                      </div>
                  </div>
                  </div>
              ))
            ) : (
              <div className='text-gray-900 font-bold header-title'>No products available</div>
            )}
            </div>
        </div>
          </div>
  );
};

export default Home;
