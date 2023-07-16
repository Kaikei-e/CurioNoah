import { SubmitHandler, useForm } from "react-hook-form";

type Props = {};

type Input = {
  searchKeyword: string;
};

export const FeedFlare = (props: Props) => {
  const {
    register,
    handleSubmit,
    watch,
    formState: { errors },
  } = useForm<Input>();

  const onSubmit: SubmitHandler<Input> = (data: Input) => console.log(data);
  return (
    <div>
      <h3>FeedFlare</h3>
      <form onSubmit={handleSubmit(onSubmit)}>
        <input defaultValue="Go" {...register("searchKeyword")} />
        <input type="submit" />
      </form>
    </div>
  );
};
