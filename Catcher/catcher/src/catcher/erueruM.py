import jax
from typing import Any, Callable, Sequence
from jax import lax, random, numpy as jnp
from flax.core import freeze, unfreeze
from flax import linen as nn

class ErErM(nn.Module):
    out_dims: int

    @nn.compact
    def __call__(self, x):
        x = x.reshape((x.shape[0], -1))
        x = nn.Dense(128)(x)  # create inline Flax Module submodules
        x = nn.relu(x)
        x = nn.Dense(self.out_dims)(x)  # shape inference
        return x


model = ErErM(out_dims=10)  # instantiate the MLP model

x = jnp.empty((4, 28, 28, 1))  # generate random data
variables = model.init(PRNGKey(42), x)  # initialize the weights
y = model.apply(variables, x)  # make forward pass